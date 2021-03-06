// Code generated by stickgen.
// DO NOT EDIT!

package repl

import (
	"fmt"
	"io"

	"github.com/tyler-sommer/stick"
)

func blockReplIndexHtmlTwigAdditionalJavascripts(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 33, offset 34 in repl/index.html.twig
	fmt.Fprint(output, `
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ace.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-searchbox.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-spellcheck.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-static_highlight.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-language_tools.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-error_marker.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/ext-whitespace.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/keybinding-vim.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/mode-javascript.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/theme-twilight.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.4.4/worker-javascript.js"></script>
<script type="text/javascript">
	$(function () {
		const $bodyField = $('#code-body');

		const editor = ace.edit('editor');
		const whitespace = ace.require('ace/ext/whitespace');
		editor.commands.addCommands(whitespace.commands);

		editor.setOptions({
			enableBasicAutocompletion: true,
			enableLiveAutocompletion: true,
			enableSnippets: false
		});
		editor.setKeyboardHandler('vim');
		editor.setTheme('ace/theme/twilight');
		editor.setValue($bodyField.val());
		editor.getSession().setMode('ace/mode/javascript');
		editor.getSession().on('change', function () {
			$bodyField.val(editor.getValue());
		});
		editor.resize();


		var $execute = $('#execute');
		var $output = $('#output');
		var $eventLog = $('#event-log');

		var es = window.squIRCyEvents;

		es.addEventListener("irc.WILDCARD", function (e) {
			var data = JSON.parse(e.data);
			$eventLog.append("[" + data.Code + "] " + data.Nick + "->" + data.Target + ": " + data.Message + "\n");
			$eventLog[0].scrollTop = $eventLog[0].scrollHeight;
		});

		$execute.on('click', function (e) {
			e.preventDefault();

			$.ajax({
				url: $execute.attr('href'),
				type: 'post',
				data: {
					script: editor.getValue(),
					scriptType: 'Javascript'
				},
				success: function (response) {
					$output.html(JSON.stringify(response, null, '  '))
				}
			});
		});
	});
</script>
`)
}
func blockReplIndexHtmlTwigContent(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 3, offset 19 in repl/index.html.twig
	fmt.Fprint(output, `
<div class="row">
	<div class="col-sm-6">
		<h4>JavaScript REPL</h4>
	</div>
	<div class="col-sm-6">
		<a class="btn btn-default btn-sm pull-right" href="https://squircy.com/js-api.html" target="_blank">Documentation <i class="fa fa-external-link"></i></a>
	</div>
</div>
<form method="post" action="/repl/execute">
<div class="row">
	<div class="col-md-8" style="height: 600px">
  		<div id="editor" style="width: 100%; height: 100%"></div>
        <div style="display: none">
            <textarea id="code-body"></textarea>
        </div>
		<br>
		<a class="form-control btn btn-primary" id="execute" href="/repl/execute">Execute</a>
	</div>
	<div class="col-md-4">
		<h5>Output</h5>
		<pre id="output" class="history" style="height: 200px">
		</pre>
		<h5>Events</h5>
		<pre id="event-log" class="history"></pre>
	</div>
</div>
</form>
`)
}

func TemplateReplIndexHtmlTwig(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
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
	blockReplIndexHtmlTwigContent(env, output, ctx)
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
	blockReplIndexHtmlTwigAdditionalJavascripts(env, output, ctx)
	// line 47, offset 18 in layout.html.twig
	fmt.Fprint(output, `
</body>

</html>
`)
	// line 1, offset 32 in repl/index.html.twig
	fmt.Fprint(output, `

`)
	// line 31, offset 14 in repl/index.html.twig
	fmt.Fprint(output, `

`)
	// line 97, offset 14 in repl/index.html.twig
	fmt.Fprint(output, `
`)
}
