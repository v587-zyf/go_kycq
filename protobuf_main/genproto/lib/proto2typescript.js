// Import in commondjs style
var DustJS = require("./dustinit");
var fs = require("fs");

// Load dust templates
loadDustTemplate("module");
loadDustTemplate("interface");
loadDustTemplate("enum");
loadDustTemplate("builder");
loadDustTemplate("service");
loadDustTemplate("define");
loadDustTemplate("serviceDef");


exports.genDeclaration = function (inputStr, options, callback) {
    // Load the json file
    var model;
    try {
        model = JSON.parse(inputStr);
    }
    catch (e) {
        callback("Input doesn't look like a JSON!", null);
    }

    // If a packagename isn't present, use a default package name
    if (!model.package) {
        model.package = "Proto2TypeScript";
    }

    // Generates the names of the model
    generateNames(options, model, model.package);

    // Render the model
    DustJS.render("module", model, function (err, out) {
        callback(err, out);
    });
};

exports.genDefinition = function(inputStr, callback) {
    var model;
    try {
        model = JSON.parse(inputStr);
    }
    catch (e) {
        callback("Input doesn't look like a JSON!", null);
    }

    var services = generateServices(model);
    delete model.services;
    var idMap = model.idMap;
    delete model.idMap;
    delete model.consts;
    var model = {
        package: model.package,
        body: JSON.stringify(model, null, 4),
        services: services,
        idMap: idMap,
    }
    DustJS.render("define", model, function(err, out) {
        callback(err, out);
    });
}

// --------------------------- functions -----------------------

function loadDustTemplate(name) {
    var template = fs.readFileSync(__dirname + "/../templates/ts/" + name + ".dust", "UTF8").toString();
    var compiledTemplate = DustJS.compile(template, name);
    DustJS.loadSource(compiledTemplate);
}

// Generate the names for the model, the types, and the interfaces
function generateNames(options, model, prefix, name) {
    if (name === void 0) {
        name = "";
    }
    model.fullPackageName = prefix + (name != "." ? name : "");

    // Copies the settings (I'm lazy)
    model.properties = options.properties;
    model.camelCaseGetSet = options.camelCaseGetSet;
    model.underscoreGetSet = options.underscoreGetSet;


    var newDefinitions = {};

    // Generate names for messages
    // Recursive call for all messages
    var key;
    for (key in model.messages) {
        var message = model.messages[key];
        newDefinitions[message.name] = "Builder";
        generateNames(options, message, model.fullPackageName, "." + (model.name ? model.name : ""));
    }

    // Generate names for enums
    for (key in model.enums) {
        var currentEnum = model.enums[key];
        newDefinitions[currentEnum.name] = "";
        currentEnum.fullPackageName = model.fullPackageName + (model.name ? "." + model.name : "");
    }

    // Generate names for consts
    for (key in model.consts) {
        var currentConst = model.consts[key];
        currentConst.fullPackageName = model.fullPackageName + (model.name ? "." + model.name : "");
    }

    // For fields of types which are defined in the same message,
    // update the field type in consequence
    for (key in model.fields) {
        var field = model.fields[key];
        if (typeof newDefinitions[field.type] !== "undefined") {
            field.type = model.name + "." + field.type;
        }
    }

    // Add the new definitions in the model for generate builders
    var definitions = [];
    for (key in newDefinitions) {
        definitions.push({name: key, type: ((model.name ? (model.name + ".") : "") + key) + newDefinitions[key]});
    }

    model.definitions = definitions;

    model.services = generateServices(model);
}

function generateServices(model) {
    var services = [];
    model.services && model.services.forEach(service => {
        for (actionName in service.rpc) {
            var action = service.rpc[actionName];
            services.push({
                name: service.name + actionName.charAt(0).toUpperCase() + actionName.slice(1),
                controller: service.name,
                action: actionName,
                request: action.request,
                response: action.response,
            });
        }
    });
    return services;
}
