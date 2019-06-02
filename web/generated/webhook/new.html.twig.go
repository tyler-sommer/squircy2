// Code generated by stickgen.
// DO NOT EDIT!

package webhook

import (
	"github.com/tyler-sommer/stick"
	"io"
	"fmt"
)

func blockWebhookNewHtmlTwigContent(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 3, offset 19 in webhook/new.html.twig
	fmt.Fprint(output, `
<div class="row">
	<div class="col-sm-6">
		<h4>New Webhook</h4>
	</div>
	<div class="col-sm-6">
		<a class="btn btn-default btn-sm pull-right" href="https://squircy.com/webhooks.html" target="_blank">Documentation <i class="fa fa-external-link"></i></a>
	</div>
</div>
<form method="post" action="/webhook/create">
<div class="row">
	<div class="col-md-7">
		<div class="row">
			<div class="col-md-5 form-group">
				<label for="hook-title">Title</label>
				<input class="form-control" id="hook-title" name="title" required placeholder="A Descriptive Title">
			</div>
			<div class="col-md-4 form-group">
				<label for="signature">Signature Header</label>
				<input class="form-control" id="signature" name="signature" required placeholder="Header name containing payload signature" value="X-Signature">
			</div>
		</div>
		<div class="row">
			<div class="col-md-5">
				<button class="form-control btn btn-primary">Create</button>
			</div>
		</div>
	</div>
	<div class="col-md-5">
		`)
	// line 1, offset 0 in webhook/_more_info.html.twig
	fmt.Fprint(output, `<div class="panel panel-primary">
    <div class="panel-heading"><i class="fa fa-question-circle"></i> More Info</div>
    <div class="panel-body">
        <p>Each incoming webhook request must contain a valid signature in the signature header defined here.</p>
        <p>A key will be automatically generated when the webhook is created, and the key is used to derive a signature unique to the payload body.</p>
        <p>A <a href="https://squircy.com/webhooks.html" target="_blank">reference Webhook sender implementation</a> is available on the squIRCy website.</p>
    </div>
</div>`)
	// line 32, offset 46 in webhook/new.html.twig
	fmt.Fprint(output, `
	</div>
</div>

</form>
`)
}
func blockWebhookNewHtmlTwigAdditionalJavascripts(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
	// line 46, offset 38 in webhook/new.html.twig
	fmt.Fprint(output, `
    `)
}

func TemplateWebhookNewHtmlTwig(env *stick.Env, output io.Writer, ctx map[string]stick.Value) {
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
	blockWebhookNewHtmlTwigContent(env, output, ctx)
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
	blockWebhookNewHtmlTwigAdditionalJavascripts(env, output, ctx)
	// line 47, offset 18 in layout.html.twig
	fmt.Fprint(output, `
</body>

</html>
`)
	// line 1, offset 32 in webhook/new.html.twig
	fmt.Fprint(output, `

`)
	// line 37, offset 14 in webhook/new.html.twig
	fmt.Fprint(output, `
`)
}
