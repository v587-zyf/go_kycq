var DustJS = require("dustjs-linkedin");
var dust = require('dustjs-helpers');
var changeCase = require('change-case');

initializeDustJS();

function initializeDustJS() {
    // Keep line breaks
    // DustJS.optimizers.format = function (ctx, node) {
    //     return node;
    // };

    // Create view filters
    DustJS.filters["firstLetterInUpperCase"] = function (value) {
        return value.charAt(0).toUpperCase() + value.slice(1);
    };

    DustJS.filters["firstLetterInLowerCase"] = function (value) {
        return value.charAt(0).toLowerCase() + value.slice(1);
    };

    DustJS.filters["lowerCase"] = function (value) {
        return changeCase.lowerCase(value);
    };

    DustJS.filters["upperCase"] = function (value) {
        return changeCase.upperCase(value);
    };

    DustJS.filters["camelCase"] = function (value) {
        return changeCase.camelCase(value);
    };

    DustJS.filters["pascalCase"] = function (value) {
        return changeCase.pascalCase(value);
    };

    DustJS.filters["constantCase"] = function (value) {
        return changeCase.constantCase(value);
    };

    DustJS.filters["convertType"] = function (value) {
        if (value.rule === 'map') {
            var type = _convertType(value.type);
            var keyType = _convertType(value.keytype);
            return "dcodeIO.ProtoBuf.Map<"+keyType+","+type+">";
        }
        return _convertType(value.type);

        function _convertType(value) {
            switch (value.toLowerCase()) {
                case 'string':
                    return 'string';
                case 'bool':
                    return 'boolean';
                case 'bytes':
                    return 'ByteBuffer';
                case 'double':
                case 'float':
                case 'int32':
                case 'int64':
                case 'uint32':
                case 'uint64':
                case 'sint32':
                case 'sint64':
                case 'fixed32':
                case 'fixed64':
                case 'sfixed32':
                case 'sfixed64':
                    return "number";
            }
            // By default, it's a message identifier
            return value;
        }
    };

    DustJS.filters["optionalFieldDeclaration"] = function (value) {
        return value == "optional" ? "?" : "";
    };

    DustJS.filters["repeatedType"] = function (value) {
        return value == "repeated" ? "[]" : "";
    };

    dust.config.whitespace = true;
    dust.helpers.iter = function(chunk, context, bodies, params) {
        var obj = dust.helpers.tap(params.obj, chunk, context);
        var type = params.type;
        var iterable = [];
        for (var key in obj) {
            if (obj.hasOwnProperty(key)) {
                var value = obj[key];
                if (type === 'keys') {
                    iterable.push(key);
                } else if (type === 'values') {
                    iterable.push(value)
                } else {
                    iterable.push({
                        '$key': key,
                        '$value': value,
                    });
                }
            }
        }
        return chunk.section(iterable, context, bodies);
    };
}

module.exports = DustJS;
