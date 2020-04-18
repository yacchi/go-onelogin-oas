import org.apache.commons.lang3.StringUtils
@Grab(group = 'org.openapitools', module = 'openapi-generator-cli', version = '4.3.0')

import org.openapitools.codegen.*
import org.openapitools.codegen.languages.*

class MyGoClientCodegen extends GoClientExperimentalCodegen {

    String name = "my-go-client-codegen"
    String modelImportPath
    String apiImportPath
    String apiFolder
    String modelsFolder

    MyGoClientCodegen() {
        super()
    }

    @Override
    String modelFileFolder() {
        return !modelsFolder ? super.modelFileFolder() : this.outputFolder + File.separator + modelsFolder + File.separator
    }

    @Override
    String apiFileFolder() {
        return !apiFolder ? super.apiFileFolder() : this.outputFolder + File.separator + apiFolder + File.separator
    }

    static List<String> toPackageName(String modulePath) {
        // import with alias style (import zzz "github.com/xxx/yyy")
        if (modulePath.contains(" ")) {
            return modulePath.split(" ", 2)
        }
        // normal import style (import "github.com/xxx/yyy")
        if (modulePath.contains("/")) {
            return [modulePath.split("/").last(), modulePath]
        }
        return [modulePath, modulePath]
    }

    @Override
    void processOpts() {
        super.processOpts()

        if (StringUtils.isNotEmpty(System.getenv("GO_PACKAGE_VERSION"))) {
            setPackageVersion(System.getenv("GO_PACKAGE_VERSION"))
            additionalProperties["packageVersion"] = packageVersion
        }

        if (additionalProperties.containsKey("modelPackage")) {
            modelPackage = (String) additionalProperties.get("modelPackage")
        }

        if (additionalProperties.containsKey("apiPackage")) {
            apiPackage = (String) additionalProperties.get("apiPackage")
        }

        (modelPackage, modelImportPath) = toPackageName(modelPackage)
        (apiPackage, apiImportPath) = toPackageName(apiPackage)

        if (modelImportPath != packageName) {
            if (modelImportPath.indexOf(packageName) == 0) {
                modelsFolder = modelImportPath.substring(packageName.length() + 1)
            }
        }
        if (apiImportPath != packageName) {
            if (apiImportPath.indexOf(packageName) == 0) {
                apiFolder = apiImportPath.substring(packageName.length() + 1)
            }
        }

        setModelPackage(modelPackage)
        setApiPackage(apiPackage)
        additionalProperties["apiPackage"] = apiPackage
        additionalProperties["apiImportPath"] = apiImportPath
        additionalProperties["modelImportPath"] = modelImportPath
        additionalProperties["modelPackage"] = modelPackage

        supportingFiles.removeIf({ o -> o.destinationFilename == "response.go" })

        for (def f : supportingFiles) {
            if (0 < f.destinationFilename.indexOf(".go")) {
                if (f.destinationFilename == "utils.go") {
                    f.folder = modelsFolder
                } else {
                    f.folder = apiFolder
                }
            }
        }
    }

    @Override
    String toModelImport(String name) {
        return "" == this.modelImportPath ? this.apiPackage : this.modelImportPath
    }

    String addPackageNamePrefix(String packageName, String name, String baseType) {
        if (this.languageSpecificPrimitives.contains(baseType)) {
            return name
        }
        return name.replace(baseType, packageName + "." + baseType)
    }

    @Override
    Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {
        if (this.modelImportPath == this.apiImportPath) {
            return super.postProcessOperationsWithModels(objs, allModels)
        }

        objs = super.postProcessOperationsWithModels(objs, allModels)

        def objectMap = (Map) objs.get("operations");
        def operations = (List<CodegenOperation>) objectMap.get("operation");

        for (def operation : operations) {
            if (operation.allParams) {
                for (def param : operation.allParams) {
                    if (param.dataType && param.baseType) {
                        param.dataType = addPackageNamePrefix(this.modelPackage, param.dataType, param.baseType)
                    }
                }
            }
            for (def res : operation.responses) {
                if (res.dataType && res.baseType) {
                    if (res.simpleType) {
                        res.baseType = toModelName(res.baseType)
                    }
                    res.dataType = addPackageNamePrefix(this.modelPackage, res.dataType, res.baseType)
                }
            }
            if (operation.returnType && operation.returnBaseType) {
                operation.returnType = addPackageNamePrefix(this.modelPackage, operation.returnType, operation.returnBaseType)
            }
        }
        return objs
    }

    static void main(String[] args) {
        OpenAPIGenerator.main(args)
    }
}
