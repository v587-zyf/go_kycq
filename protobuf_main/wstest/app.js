var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var _ = require('lodash');
var pb = require('./public/js/proto.json');
var config = require('./config.json');
var allSchema = require('./public/js/protoschema.json')

var app = express();

var env = process.env.NODE_ENV || 'development';
app.locals.ENV = env;
app.locals.ENV_DEVELOPMENT = env == 'development';

// view engine setup

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

// app.use(favicon(__dirname + '/public/img/favicon.ico'));
app.use(express.static(path.join(__dirname, 'public')));

function getSchema(msg) {
    return allSchema[msg] && allSchema[msg]['definitions'] && allSchema[msg]['definitions']['pb.' + msg] || {}
}

app.use('/info', function (req, res, next) {
    var messages = _.keyBy(pb.messages, 'name');
    var request, respose;

    if (req.query.service == 'commands') {
        request = req.query.action;
        response = '';
    } else {
        var service = _.find(pb.services, obj => { return obj.name == req.query.service });
        var action = service.rpc[req.query.action];
        request = action.request;
        response = action.response;
    }

    res.render("jsoninfo", {
        request: request,
        response: response,
        messages: messages,
        schema: getSchema(request)
    });
});

app.use('/', function (req, res, next) {
    var categories = {};
	commands = _.values(pb.idMap, 'name');
	var newCommands = []
	var reg=/Req$/;
	for ( i in commands )
	{
		if (reg.test(commands[i]))
		{
			newCommands.push(commands[i])
		}
	}
    categories['commands'] = newCommands;
    _.each(pb.services, service => categories[service.name] = _.keys(service.rpc));
    res.render('index', {
        title: config.name,
        format: config.format,
        messages: _.keyBy(pb.messages, 'name'),
        services: _.keyBy(pb.services, 'name'),
        categories: categories,
        config: config,
    });
});

/// catch 404 and forward to error handler
app.use(function (req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});

/// error handlers

// development error handler
// will print stacktrace

if (app.get('env') === 'development') {
    app.use(function (err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
            message: err.message,
            error: err,
            title: 'error'
        });
    });
}

// production error handler
// no stacktraces leaked to user
app.use(function (err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
        message: err.message,
        error: {},
        title: 'error'
    });
});


var port = config.port || 3501;

var server = app.listen(port, function () {
    console.log('Express server listening on port ' + server.address().port);
});


