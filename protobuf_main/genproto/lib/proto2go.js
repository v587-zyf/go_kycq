// Import in commondjs style
var DustJS = require("./dustinit");
var fs = require("fs");

// Load dust templates
loadDustTemplate("idmap");
loadDustTemplate("const");

exports.genIdMap = function(model, callback) {
    // Render the model
    DustJS.render("idmap", model, function(err, out) {
        callback(err, out);
    });
};

exports.genEnums = function(model, callback) {
    DustJS.render("const", model, function(err, out) {
        callback(err, out)
    });
}
// --------------------------- functions -----------------------

function loadDustTemplate(name) {
    var template = fs.readFileSync(__dirname + "/../templates/go/" + name + ".dust", "UTF8").toString();
    var compiledTemplate = DustJS.compile(template, name);
    DustJS.loadSource(compiledTemplate);
}
