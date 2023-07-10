#!/usr/bin/env node

var fs = require('fs-extra');
var path = require('path');
var pbjs = require('protobufjs/cli/pbjs')
var globArray = require('glob-array');
var proto2ts = require('./lib/proto2typescript');
var proto2go = require('./lib/proto2go');
var DustJS = require("dustjs-linkedin");
var idmapParser = require('./lib/idmap/index');
var argv = require('yargs')
        .help('h')
        .option('file', {alias: 'f', describe: '指定协议文件', demand: true})
        .option('gen', {alias: 'g', describe: '生成对应语言的文件', choices: ['ts', 'go', 'json', 'jsonschema'], demand: true})
        .option('out', {alias: 'o', describe: '输出目录', default: './gen'})
        .argv;

function genproto(fileName) {
    var info = idmapParser(fileName);
    if (argv.gen == 'go') {
        outputGo({packageName: info.packageName, idMap: info.idMap, logLvMap: info.logLvMap, enums: info.enums});
        return;
    }

    var pbFiles = info.pbFiles;
    if (typeof pbFiles === 'string') pbFiles = [pbFiles];
    pbFiles = pbFiles.map(file => path.join(path.dirname(fileName), file))

    var files = globArray.sync(pbFiles, null);
    var includePath = [];
    files.forEach(file => {
        var p = path.resolve(path.dirname(file));
        if (includePath.indexOf(p) < 0)
            includePath.push(p);
    });
    var options = {
        path: includePath,
    };
    var builder = pbjs.sources['proto'](files, options)
    var result = pbjs.targets['json'](builder, options);
    result = JSON.parse(result);
    result.idMap = info.idMap;
    //没有找到在哪里可以设置 ，这里强制设置成proto3
    result.syntax = "proto3"

    if (argv.gen == 'jsonschema') {
		outputJsonSchema(files, info.logLvMap);
    } else if (argv.gen == 'json') {
        result.reqMap = info.reqMap
        fs.outputFileSync(path.join(argv.out, 'proto.json'), JSON.stringify(result, null, 4));
    } else {
        result.consts = info.enums
        outputTs(JSON.stringify(result))
    }
}

function outputTs(source) {
    proto2ts.genDeclaration(source, {properties: true}, (err, out) => {
        if (err) return console.log(err);
        fs.outputFileSync(path.join(argv.out, 'proto.d.ts'), out);
    });
    proto2ts.genDefinition(source, (err, out) => {
        if (err) return console.log(err);
        fs.outputFileSync(path.join(argv.out, 'proto.js'), out);
    });
}

function outputGo(source) {
    proto2go.genIdMap(source, (err, out) => {
        if (err) return console.log(err);
        fs.outputFileSync(path.join(argv.out, 'idmap.go'), out);
    });
    proto2go.genEnums(source, (err, out) => {
        if (err) return console.log(err);
        fs.outputFileSync(path.join(argv.out, 'const.go'), out);
    });
}

function outputJsonSchema(pbFiles, reqMap){
    var protobuf2jsonschema = require("protobuf-jsonschema")
    var allObj = {};
    // 由于protobuf-jsonschema不能正确处理map<int32, int32>类型，故这里每个Message独自解析
    pbFiles.forEach(file => {
        for (var k in reqMap) {
            try {
                var obj = protobuf2jsonschema(file, 'pb.' + k);
                allObj[k] = obj;
            } catch(err) {
				//console.log(err)
            }
        }
    });

    var writeFilePath = path.join(argv.out,'protoschema.json')
    fs.outputFileSync(writeFilePath, JSON.stringify(allObj, null, 4));
}

module.exports = genproto

if (require.main == module) {
    var file = path.isAbsolute(argv.file) ? argv.file : path.join(process.cwd(), argv.file);
    if (!fs.existsSync(file)) {
        console.log('file not exist');
        return;
    }
    if (!fs.existsSync(argv.out))
        fs.ensureDir(argv.out);
    genproto(file);
}

