// Code generated by stickgen.
// DO NOT EDIT!

package manage

import (
	"github.com/tyler-sommer/stick"
	"io"
	"fmt"
)

func blockManageEditHtmlTwigContent(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 3, offset 19 in manage/edit.html.twig
	fmt.Fprint(output, `
<div class="row">
	<div class="col-sm-6">
		<h4>Modify Settings</h4>
	</div>
	<div class="col-sm-6">
		<a class="btn btn-default btn-sm pull-right" href="https://squircy.com" target="_blank">Documentation <i class="fa fa-external-link"></i></a>
	</div>
</div>
<form method="post" action="/manage/update">
<div class="row">
	<div class="col-md-4">
		<h5>IRC</h5>
		<div class="form-group">
			<label for="network">Network</label>
			<input class="form-control" name="network" id="network" placeholder="irc.freenode.net:6667" value="`)
	// line 18, offset 102 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "Network")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 18, offset 122 in manage/edit.html.twig
	fmt.Fprint(output, `">
			<div class="checkbox">
				<label for="tls">
					<input type="checkbox" name="tls" id="tls" value="on" `)
	// line 21, offset 62 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "TLS")
		if err == nil && stick.CoerceBool(val) {
			// line 21, offset 78 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 21, offset 97 in manage/edit.html.twig
	fmt.Fprint(output, `> Use TLS
				</label>
			</div>
			<div class="checkbox">
				<label for="auto_connect">
					<input type="checkbox" name="auto_connect" id="auto_connect" value="on" `)
	// line 26, offset 80 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "AutoConnect")
		if err == nil && stick.CoerceBool(val) {
			// line 26, offset 104 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 26, offset 123 in manage/edit.html.twig
	fmt.Fprint(output, `> Connect to IRC when squIRCy2 starts
				</label>
			</div>
		</div>
	</div>
	<div class="col-md-6 col-md-offset-2">
		`)
	// line 1, offset 0 in manage/_save_button.html.twig
	fmt.Fprint(output, `<div class="panel panel-info">
    <div class="panel-heading">
        <i class="fa fa-exclamation-triangle"></i> Most changes require fully restarting the bot before they will take effect.
    </div>
    <div class="panel-body">
        <button class="btn btn-primary">Save Changes</button>
    </div>
</div>`)
	// line 32, offset 47 in manage/edit.html.twig
	fmt.Fprint(output, `
	</div>
</div>
<div class="row">
	<div class="col-md-4">
		<div class="form-group">
			<label for="nick">Bot Nickname</label>
			<input class="form-control" name="nick" id="nick" placeholder="mrsquishy" value="`)
	// line 39, offset 84 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "Nick")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 39, offset 101 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>
		<div class="form-group">
			<label for="username">Bot Username</label>
			<input class="form-control" name="username" id="username" placeholder="mrjones" value="`)
	// line 43, offset 90 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "Username")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 43, offset 111 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>
		<div class="form-group">
			<label for="owner_nick">Owner Nickname</label>
			<input class="form-control" name="owner_nick" id="owner_nick" placeholder="" value="`)
	// line 47, offset 87 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "OwnerNick")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 47, offset 109 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>
		<div class="form-group">
			<label for="owner_host">Owner Hostname</label>
			<input class="form-control" name="owner_host" id="owner_host" placeholder="" value="`)
	// line 51, offset 87 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "OwnerHost")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 51, offset 109 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>
	</div>
	<div class="col-md-4 col-md-offset-1">
		<div class="form-group">
			<div class="checkbox">
				<label for="enable_sasl">
					<input type="checkbox" name="enable_sasl" id="enable_sasl" value="on" `)
	// line 58, offset 78 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SASL")
		if err == nil && stick.CoerceBool(val) {
			// line 58, offset 95 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 58, offset 114 in manage/edit.html.twig
	fmt.Fprint(output, `> Enable SASL Authentication
				</label>
			</div>
		</div>
		<div class="form-group">
			<label for="sasl_nickname">SASL Username</label>
			<input class="form-control" name="sasl_username" id="sasl_username" placeholder="" value="`)
	// line 64, offset 93 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SASLUsername")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 64, offset 118 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>
		<div class="form-group">
			<label for="sasl_password">SASL Password</label>
			<input class="form-control" type="password" name="sasl_password" id="sasl_password" value="`)
	// line 68, offset 94 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SASLPassword")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 68, offset 119 in manage/edit.html.twig
	fmt.Fprint(output, `">
		</div>

	</div>
</div>

	<div class="row">
		<div class="col-md-4">
			<h5>Script Management</h5>
			<div class="form-group">
				<div class="checkbox">
					<label for="scripts_as_files">
						<input type="checkbox" name="scripts_as_files" id="scripts_as_files" value="on" `)
	// line 80, offset 89 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "ScriptsAsFiles")
		if err == nil && stick.CoerceBool(val) {
			// line 80, offset 116 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 80, offset 135 in manage/edit.html.twig
	fmt.Fprint(output, `> Store Scripts as regular files on the Filesystem
					</label>
				</div>
			</div>
			<div class="form-group">
				<label for="scripts_path">Scripts Storage Path</label>
				<input class="form-control" name="scripts_path" id="scripts_path" placeholder="`)
	// line 86, offset 83 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "RootPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 86, offset 104 in manage/edit.html.twig
	fmt.Fprint(output, `/scripts/" value="`)
	// line 86, offset 122 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "ScriptsPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 86, offset 146 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
			<div class="form-group">
				<label>Script Utilities</label><br>
				<div class="btn-group">
					<a id="export-scripts" class="btn btn-default"><i class="fa fa-external-link"></i> Export to Filesystem</a>
					<a id="import-scripts" class="btn btn-default btn-default-alt"><i class="fa fa-external-link fa-rotate-180"></i> Import from Filesystem</a>
				</div>
			</div>
			<div class="form-group">
				<div class="checkbox">
					<label for="enable_file_api">
						<input type="checkbox" name="enable_file_api" id="enable_file_api" value="on" `)
	// line 98, offset 87 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "EnableFileAPI")
		if err == nil && stick.CoerceBool(val) {
			// line 98, offset 113 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 98, offset 132 in manage/edit.html.twig
	fmt.Fprint(output, `> Enable Filesystem API in Scripts <a style="font-size: 1.1em" href="https://squircy.com/js-api.html#file-api" target="_blank"><i class="fa fa-question-circle"></i></a>
					</label>
				</div>
			</div>
			<div class="form-group">
				<label for="file_api_root">File API Root Path</label>
				<input class="form-control" name="file_api_root" id="file_api_root" placeholder="`)
	// line 104, offset 85 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "RootPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 104, offset 106 in manage/edit.html.twig
	fmt.Fprint(output, `/storage/" value="`)
	// line 104, offset 124 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "FileAPIRoot")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 104, offset 148 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
		</div>
		<div class="col-md-4 col-md-offset-1">
			<h5>Plugins</h5>
			<div class="form-group">
				<label for="plugins_enabled">Enabled Plugins</label>
				<input class="form-control" name="plugins_enabled" id="plugins_enabled" placeholder="plugin1,plugin2" value="`)
	// line 111, offset 113 in manage/edit.html.twig
	{
		val, _ := stick.GetAttr(ctx["config"], "PluginsEnabled")

		var fnval stick.Value = ""
		if fn, ok := env.Filters["join"]; ok {
			fnval = fn(nil, val, ",", )
		}
		fmt.Fprint(output, fnval)
	}
	// line 111, offset 150 in manage/edit.html.twig
	fmt.Fprint(output, `">
				<p>Comma separated list without the .so extension</p>
			</div>
			<div class="form-group">
				<label for="plugins_path">Plugin Root Path</label>
				<input class="form-control" name="plugins_path" id="plugins_path" placeholder="/path/to/squircy/plugins" value="`)
	// line 116, offset 116 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "PluginsPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 116, offset 140 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
		</div>
	</div>
	<div class="row">
		<div class="col-md-4">
			<h5>Web Server</h5>
			<div class="form-group">
				<div class="checkbox">
					<label for="web_interface">
						<input type="checkbox" name="web_interface" id="web_interface" value="on" `)
	// line 126, offset 83 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "WebInterface")
		if err == nil && stick.CoerceBool(val) {
			// line 126, offset 108 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 126, offset 127 in manage/edit.html.twig
	fmt.Fprint(output, `> Enable Web Interface
					</label>
				</div>
				<label for="http_host_port">HTTP Listen Host and Port</label>
				<input class="form-control" name="http_host_port" id="http_host_port" placeholder=":3000" value="`)
	// line 130, offset 101 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "HTTPHostPort")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 130, offset 126 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
			<br>
			<h5>HTTPS</h5>
			<div class="form-group">
				<div class="checkbox">
					<label for="https">
						<input type="checkbox" name="https" id="https" value="on" `)
	// line 137, offset 67 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "HTTPS")
		if err == nil && stick.CoerceBool(val) {
			// line 137, offset 85 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 137, offset 104 in manage/edit.html.twig
	fmt.Fprint(output, `> Enable HTTPS
					</label>
				</div>
				<div class="checkbox">
					<label for="require_https">
						<input type="checkbox" name="require_https" id="require_https" value="on" `)
	// line 142, offset 83 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "RequireHTTPS")
		if err == nil && stick.CoerceBool(val) {
			// line 142, offset 108 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 142, offset 127 in manage/edit.html.twig
	fmt.Fprint(output, `> Require HTTPS
					</label>
				</div>
			</div>
		</div>
		<div class="col-md-4 col-md-offset-1">
			<div class="form-group">
				<h5>HTTP(S) Authentication</h5>
				<div class="checkbox">
					<label for="http_auth">
						<input type="checkbox" name="http_auth" id="http_auth" value="on" `)
	// line 152, offset 75 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "HTTPAuth")
		if err == nil && stick.CoerceBool(val) {
			// line 152, offset 96 in manage/edit.html.twig
			fmt.Fprint(output, `checked `)
		}
	}
	// line 152, offset 115 in manage/edit.html.twig
	fmt.Fprint(output, `> Enable HTTP BasicAuth
					</label>
				</div>
			</div>
			<div class="form-group">
				<label for="auth_username">Username</label>
				<input class="form-control" name="auth_username" id="auth_username" placeholder="Username" value="`)
	// line 158, offset 102 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "AuthUsername")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 158, offset 127 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
			<div class="form-group">
				<label for="auth_password">Password</label>
				<input class="form-control" name="auth_password" id="auth_password" placeholder="Password" value="`)
	// line 162, offset 102 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "AuthPassword")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 162, offset 127 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
		</div>
	</div>
	<div class="row">
		<div class="col-md-4">
			<div class="form-group">
				<label for="ssl_host_port">SSL Listen Host and Port</label>
				<input class="form-control" name="ssl_host_port" id="ssl_host_port" placeholder=":3443" value="`)
	// line 170, offset 99 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SSLHostPort")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 170, offset 123 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
			<div class="form-group">
				<label for="ssl_cert_file">SSL Certificate File</label>
				<input class="form-control" name="ssl_cert_file" id="ssl_cert_file" placeholder="/path/to/cert.pem" value="`)
	// line 174, offset 111 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SSLCertFile")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 174, offset 135 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
		</div>
	</div>
<div class="row">
	<div class="col-md-4">
		<div class="row">
			<div class="col-md-12 form-group">
				<label for="ssl_cert_key">SSL Certificate Key File</label>
				<input class="form-control" name="ssl_cert_key" id="ssl_cert_key" placeholder="/path/to/key.pem" value="`)
	// line 183, offset 108 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "SSLCertKey")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 183, offset 131 in manage/edit.html.twig
	fmt.Fprint(output, `">
			</div>
		</div>
	</div>
	<div class="col-md-6 col-md-offset-2">
		`)
	// line 1, offset 0 in manage/_save_button.html.twig
	fmt.Fprint(output, `<div class="panel panel-info">
    <div class="panel-heading">
        <i class="fa fa-exclamation-triangle"></i> Most changes require fully restarting the bot before they will take effect.
    </div>
    <div class="panel-body">
        <button class="btn btn-primary">Save Changes</button>
    </div>
</div>`)
	// line 188, offset 47 in manage/edit.html.twig
	fmt.Fprint(output, `
	</div>
</div>
<div class="row">

</div>
</form>
<form action="/manage/import-scripts" method="post">
<div id="confirm-import" class="modal fade" tabindex="-1" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
				<h4 class="modal-title">Import Scripts from Filesystem</h4>
			</div>
			<div class="modal-body">
				<p>Continue importing scripts from filesystem?</p>
				<div class="alert alert-info"><strong><i class="fa fa-info-circle"></i> Note —</strong> Existing scripts will be updated.</div>
				<div class="form-group">
					<label for="scripts_import_path">Scripts stored in squIRCy2's database will be imported from:</label>
					<input class="form-control" name="scripts_import_path" id="scripts_import_path" value="`)
	// line 208, offset 92 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "ScriptsPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 208, offset 116 in manage/edit.html.twig
	fmt.Fprint(output, `">
				</div>
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
				<button type="submit" class="btn btn-primary">Import!</button>
			</div>
		</div>
	</div>
</div>
</form>
<form action="/manage/export-scripts" method="post">
<div id="confirm-export" class="modal fade" tabindex="-1" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
				<h4 class="modal-title">Export Scripts to Filesystem</h4>
			</div>
			<div class="modal-body">
				<p>Continue exporting scripts to filesystem?</p>
				<div class="alert alert-info">
					<p><strong><i class="fa fa-info-circle"></i> Note —</strong> Existing scripts will be updated.</p>
				</div>
				<div class="form-group">
					<label for="scripts_export_path">Scripts stored in squIRCy2's database will be exported to:</label>
					<input class="form-control" name="scripts_export_path" id="scripts_export_path" value="`)
	// line 234, offset 92 in manage/edit.html.twig
	{
		val, err := stick.GetAttr(ctx["config"], "ScriptsPath")
		if err == nil {
			fmt.Fprint(output, val)
		}
	}
	// line 234, offset 116 in manage/edit.html.twig
	fmt.Fprint(output, `">
				</div>
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
				<button type="submit" class="btn btn-primary">Export!</button>
			</div>
		</div>
	</div>
</div>
</form>
`)
}
func blockManageEditHtmlTwigAdditionalJavascripts(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 247, offset 34 in manage/edit.html.twig
	fmt.Fprint(output, `
<script type="text/javascript">
	$(function() {
	    $('#import-scripts').on('click', function(e) {
	        e.preventDefault();
	        e.stopPropagation();
	        $('#confirm-import').modal('show');
		});

        $('#export-scripts').on('click', function(e) {
            e.preventDefault();
            e.stopPropagation();
            $('#confirm-export').modal('show');
        });
	})
</script>
`)
}

func TemplateManageEditHtmlTwig(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 1, offset 0 in layout.html.twig
	fmt.Fprint(output, `<!DOCTYPE html>
<html>
<head>
  <title>squIRCy</title>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/moment.js/2.24.0/moment.min.js"></script>
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/bootswatch/3.4.0/cyborg/bootstrap.min.css">
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
  <link rel="stylesheet" href="/css/style.css">
</head>

<body>
	<div id="main-container" class="container">
		<div class="row">
			<div class="col-md-12">
			`)
	// line 17, offset 6 in layout.html.twig
	blockManageEditHtmlTwigContent(env, output, ctx)
	// line 18, offset 17 in layout.html.twig
	fmt.Fprint(output, `
			</div>
		</div>
	</div>

	<nav id="main-nav" class="navbar navbar-inverse navbar-fixed-top" role="navigation">
	  	<div class="container">
			<div class="navbar-header">
				<a class="navbar-brand" href="https://squircy.com" target="_blank">squIRCy</a>
        	</div>
			<ul class="nav navbar-nav">
				<li><a href="/">Dashboard</a></li>
				<li><a href="/script">Scripts</a></li>
				<li><a href="/webhook">Webhooks</a></li>
				<li><a href="/manage">Settings</a></li>
				<li class="divider">&nbsp;</li>
				<li><a href="/repl">REPL</a></li>
			</ul>
			<ul class="nav navbar-nav pull-right">
				<li><a id="reinit" title="Re-initialize scripts" href="/script/reinit"><i class="fa fa-refresh"></i></a></li>
		        <li><a class="post-action" id="connect-button" style="display: none" href="/connect"><i class="fa fa-power-off"></i> Connect</a></li>
				<li><a class="post-action" id="disconnect-button" style="display: none" href="/disconnect"><i class="fa fa-power-off"></i> Disconnect</a></li>
				<li><a class="post-action" id="connecting-button" style="display: none" href="/disconnect"><i class="fa fa-power-off"></i> Connecting...</a></li>
		      </ul>
	  	</div>
	</nav>

    `)
	// line 1, offset 0 in _page_javascripts.html.twig
	fmt.Fprint(output, `<script type="text/javascript">
$(function() {
    var es = window.squIRCyEvents = new EventSource('/event');

	var $connectBtn = $('#connect-button');
	var $disconnectBtn = $('#disconnect-button');
	var $connectingBtn = $('#connecting-button');
    es.addEventListener("irc.CONNECTING", function() {
        $connectingBtn.show();
		$disconnectBtn.add($connectBtn).hide();
    });
    es.addEventListener("irc.CONNECT", function() {
        $disconnectBtn.show();
		$connectBtn.add($connectingBtn).hide();
    });
    es.addEventListener("irc.DISCONNECT", function() {
        $connectBtn.show();
        $disconnectBtn.add($connectingBtn).hide();
    });

	var $postLinks = $('.post-action');
	$postLinks.on('click', function(e) {
		e.preventDefault();
		
		var url = $(this).attr('href');
		$.ajax({
			url: url,
			type: 'post'
		});
	});
	
	var $reinit = $('#reinit').tooltip({
		placement: 'bottom',
		container: 'body'
	});
	$reinit.on('click', function(e) {
		e.preventDefault();
		
		if (confirm('Are you sure you want to reinitialize all scripts?')) {
			var url = $(this).attr('href');
			$.ajax({
				url: url,
				type: 'post'
			})
		}
	});

	$.ajax({
		url: '/status',
		type: 'get',
		dataType: 'json',
		success: function(response) {
			if (response.Status == 2) {
				$disconnectBtn.show();
				$connectBtn.add($connectingBtn).hide();
			} else if (response.Status == 1) {
				$connectingBtn.show();
				$disconnectBtn.add($connectBtn).hide();
			} else {
				$connectBtn.show();
				$disconnectBtn.add($connectingBtn).hide();
			}
		}
	});
});
</script>`)
	// line 45, offset 47 in layout.html.twig
	fmt.Fprint(output, `
    `)
	// line 46, offset 7 in layout.html.twig
	blockManageEditHtmlTwigAdditionalJavascripts(env, output, ctx)
	// line 47, offset 18 in layout.html.twig
	fmt.Fprint(output, `
</body>

</html>
`)
	// line 1, offset 32 in manage/edit.html.twig
	fmt.Fprint(output, `

`)
	// line 245, offset 14 in manage/edit.html.twig
	fmt.Fprint(output, `

`)
}
