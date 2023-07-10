const fs = require('fs');
const path = require('path');
const jison = require('jison');
const bnf = fs.readFileSync(path.join(__dirname, 'idmap.jison'), "utf8");
var parser = new jison.Parser(bnf);
parser.yy = {
    setPackageName: setPackageName,
    addIdMap: addIdMap,
    setPbFiles: setPbFiles,
    includeFile: includeFile,
    addEnum: addEnum,
};

var basePath = '';
var idMap = {};
var logLvMap = {};
var enums = [];
var pbFiles = '';
var packageName = '';

function setPackageName(name) {
    packageName = name;
}

function addIdMap(id, name, logLv) {
    idMap[id] = name;
    logLvMap[name] = logLv;
}

function setPbFiles(files) {
    pbFiles = files;
}

function addEnum(enumName, fields) {
    enums.push({name: enumName, values: fields});
}

function includeFile(fileName) {
    fileName = path.normalize(path.join(basePath, fileName));
    parser.parse(fs.readFileSync(fileName, 'utf-8'));
}

function parse(fileName) {
    basePath = path.normalize(path.dirname(fileName));
    parser.parse(fs.readFileSync(fileName, 'utf-8'));
    return {packageName: packageName, idMap: idMap, logLvMap: logLvMap, pbFiles: pbFiles, enums: enums};
}

module.exports = parse;

if (require.main == module) {
    if (process.argv.length < 3) {
        console.log('please specify a file');
        return;
    }
    var fileName = process.argv[2];
    var fileFullName = path.join(process.cwd(), fileName);
    if (!fs.existsSync(fileName)) {
        console.log(fileName + ' does not exist');
        return;
    }
    load(fileFullName);
}
