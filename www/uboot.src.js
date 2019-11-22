// <script type="text/javascript">if (typeof jQuery == 'undefined') document.writeln("<script src=\"uboot.js?ver="+new Date().getTime()+"\" type=\"text/javascript\"></" + "script>");</script>
var bootPATH = __CreateJSPath("uboot.js");
function __CreateJSPath(js) {
	var scripts = document.getElementsByTagName("script");
	var path = "";
	for (var i = 0, l = scripts.length; i < l; i++) {
		var src = scripts[i].src;
		if (src.indexOf(js) != -1) {
			var ss = src.split(js);
			path = ss[0];
			break;
		}
	}
	var href = location.href;
	href = href.split("#")[0];
	href = href.split("?")[0];
	var ss = href.split("/");
	ss.length = ss.length - 1;
	href = ss.join("/");
	if (path.indexOf("https:") == -1 && path.indexOf("http:") == -1 && path.indexOf("file:") == -1 && path.indexOf("\/") != 0) {
		path = href + "/" + path;
	}
	return path;
}

document.writeln(`<link 		href="//cdn.bootcss.com/monaco-editor/0.15.6/min/vs/editor/editor.main.css" data-name="//cdn.bootcss.com/monaco-editor/0.15.6/min/vs/editor/editor.main" rel="stylesheet">`);

document.writeln(`<script 		src ="${bootPATH}/scripts/jquery.min.js" type="text/javascript"></sc${'ript'}>`);
document.writeln(`<link 		href="${bootPATH}/scripts/miniui/themes/default/miniui.css" rel="stylesheet" type="text/css" />`);
document.writeln(`<script 		src ="${bootPATH}/scripts/miniui/miniui.js" type="text/javascript"></sc${'ript'}>`);
document.writeln(`<link 		href="${bootPATH}/scripts/miniui/themes/icons.css" rel="stylesheet" type="text/css" />`);
document.writeln(`<script 		src ="${bootPATH}/scripts/vue.js" type="text/javascript"></sc${'ript'}>`);
document.writeln(`<script 		>Vue.config.productionTip = false;</sc${'ript'}>`);
document.writeln(`<link  		href="${bootPATH}/scripts/iview.css" rel="stylesheet" type="text/css" />`);
document.writeln(`<script 		src ="${bootPATH}/scripts/iview.min.js" type="text/javascript"></sc${'ript'}>`);

document.writeln(`<script>var require = { paths: { 'vs': '//cdn.bootcss.com/monaco-editor/0.15.6/min/vs' } };</sc${'ript'}>`);
document.writeln(`<script defer src ="//cdn.bootcss.com/monaco-editor/0.15.6/min/vs/loader.js"></sc${'ript'}>`);
document.writeln(`<script defer src ="//cdn.bootcss.com/monaco-editor/0.15.6/min/vs/editor/editor.main.nls.js"></sc${'ript'}>`);
document.writeln(`<script defer src ="//cdn.bootcss.com/monaco-editor/0.15.6/min/vs/editor/editor.main.js"></sc${'ript'}>`);

document.writeln(`<script src="${bootPATH}scripts/core.js" type="text/javascript"></sc${'ript'}>`);
document.writeln(`<script src="${bootPATH}scripts/algorithms/base64.js" type="text/javascript"></sc${'ript'}>`);

document.writeln(`<script 		src ="//cdn.bootcss.com/highlight.js/9.15.10/highlight.min.js"></sc${'ript'}>`);
document.writeln(`<link  		href="//cdn.bootcss.com/highlight.js/9.15.10/styles/vs2015.min.css" rel="stylesheet" type="text/css" />`);


var Toast={
	info:function(msg,interval=2){
		app.$Message.info({content: msg,duration: interval});
	}
	,success:function(msg,interval=2){
		app.$Message.success({content: msg,duration: interval});
	}
	,warning:function(msg,interval=2){
		app.$Message.warning({content: msg,duration: interval});
	}
	,error:function(msg,interval=2){
		app.$Message.error({content: msg,duration: interval});
	}
}

function GetE(e){
	return e ? e : window.event;
}
function ParseQueryString(val) {
	var uri = window.location.search;
	var re = new RegExp("" + val + "=([^&?]*)", "ig");
	return ((uri.match(re)) ? (uri.match(re)[0].substr(val.length + 1)) : null);
}

function FormValid(id){
	let form=null;
	try{
		form = new mini.Form(`#${id}`);
		form.validate();
		
	}catch(ex){
	}finally{
		if(null!=form)return false==form.isValid()?null:form;
		return form;
	}
}
function DialogForm(id){
	return new mini.Form($(`#${id} .form`).attr('id'));
}
function DialogShow(id,title,done,row){
	//	<Dialog id="ReportInfoWin" class="mini-window" 
	//		<table v-bind:id="Math.guid()" class="form">
	mini.get(`${id}`).set({ "title": title}).show();
	$(`#${id} .submit`).attr('onclick', `${done}`);
	var form = new mini.Form($(`#${id} .form`).attr('id'));
	form.clear();
	form.loading();
	if(null!=row)form.setData(row);
	form.unmask();
	return form;
}
function DialogHide(id){
	mini.get(id).hide();
}


function WashGridRows(rows){
	rows.forEach(function(a,b,c){
		delete a["_id"];
		delete a["_uid"];
		delete a["_state"];
		c[b]=a;
	});
	return rows;
}

function PagingGrid(id,e){
	var pageIndex	=0;
	var pageSize	=mini.get(id).pageSize;
	if(e){
		e.cancel = true;
		if(e.data){
			pageIndex	=e.data.pageIndex;
			pageSize	=e.data.pageSize;
		}
	}
	return {
		pageIndex:pageIndex,
		pageSize:pageSize,
		id:id,
	};
}
function ClearGrid(id,col,row){
	with(mini.get(id)){
		if(col||true)setColumns([]);
		if(row||true)setData([]);
	}
}
function FillGrid(id,resp){
	if(!id||id.length<1)return;
	if(null==resp||!resp.data)return;
	let grid=mini.get(id);
	with(resp){
		grid.setData(data);
		grid.setTotalCount(totalCount||data.length);
		grid.setPageSize(size);
		grid.setPageIndex(page);
	}
}


function ClickLayout(id,layout){//west,east,south,north
	$(`#${id} > div > div.mini-layout-proxy.mini-layout-proxy-${layout}`).trigger("click")
}


// parse=function(text,index,{title,field,order})
//
// let cols=GenGridColumn(row,function(t,i,c){
// 	if(-1!=t.indexOf("@")){
// 		let parts = t.split("@");
// 		c.title=parts[0];
// 		c.order=parseInt(parts[1]);
// 	}
// 	return c;
// })
function GenGridColumn(row,parse,orderAsc){
	var cols=[];
	var keys=Object.getOwnPropertyNames(row);
	keys.forEach(function(a,b,c){
		// let title=a;
		// let field=a;
		// let order=cols.length;
		// if(-1!=a.indexOf("@")){
		// 		let parts = a.split("@");
		// 		title=parts[0];
		// 		order=parseInt(parts[1]);
		// }
		let col={"title":a,"field":a,"order":cols.length};
		if(parse)col=parse(a,b,col);
		with(col){
			let val=row[a];
			val=val||'';
			switch(typeof(val)){
				case "number":
					cols.push({field: field, 	header: title,order:order, 	renderer: ''						,	_editor: { type: 'spinner' 			},															allowSort: true, 	width: 60, 						align: 'center', 	headerAlign: 'center', 	})
					break;
				case "boolean":
					cols.push({field: field, 	header: title,order:order,	_renderer: ''						,/* editor: {*/type: 'checkboxcolumn',/*},*/														/*allowSort: true,*/ width: 80,						align: 'center', 	headerAlign: 'center',	})
					break;
				case "string":
				case "date":
					if(-1!=val.indexOf("T")&&-1!=val.indexOf("Z")&&val.endsWith("Z")){
					cols.push({field: field, 	header: title,order:order,	renderer: "app.iso8601Renderer"		,	_editor: { type: 'datepicker' 		},															allowSort: true, 	width: 180, 					align: 'center', 	headerAlign: 'center',	})
					}else{
					cols.push({field: field, 	header: title,order:order,	_renderer: ''						,	_editor: { type: 'textbox' 			},															allowSort: true, 	width: 120+20*title.length, 	align: 'center', 	headerAlign: 'center', 	})
					}
					break;
				// case "enum":
				// 	break;
				// case "multistring":
				// 	// 	renderer=`<textarea name="${key}" style="width:${formVal('width')}px;height:80px;" class="mini-textarea" emptyText="请输入"></textarea>`;
				// 	break;
				default:
					alert(`unknow type[${type}]`);
					debugger
					break;
			}
		}
	});
	cols.push({field: "", 		header: "",order:cols.length,		_renderer: ''						,	allowSort: false, 	width: '100%', 	align: 'center', 	headerAlign: 'center', 	})
	cols.sort(function(a,b){
		return (orderAsc||true)?a.order>b.order:a.order<b.order;
	});
	return cols;
}
function HasSelect(id,msg,callback){
	var grid = mini.get(id);
	var rows = grid.getSelecteds();
	if (rows.length < 1) {
		if(null!=msg)app.$Message.error({content: msg,duration: 2});
		return;
	}
	callback(rows);
}
function Confirm(msg,callback){
	//alert('confirm');
	mini.confirm(msg, "询问：",function (action) {
		if (action !== "ok") return;
		callback();
	});
}

PostData=function(route,params,ok,silence,done){
	silence=silence||false;
	if(!silence)mini.mask({el: document.body,cls: 'mini-mask-loading',html: '\u64cd\u4f5c\u4e2d\u002e\u002e\u002e'});
	setTimeout(function(){
		$.ajax({
			type: "POST",
			url: route,
			contentType: "application/json; charset=utf-8",
			data: JSON.stringify(params),
			dataType: "json",
			beforeSend: function(xhr) {
				xhr.setRequestHeader("Authorization", "Bearer " + getCookie('jwt'));
				//xhr.setRequestHeader("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI3SVJUX3N6bkVkc3hIZUl0V3NCZmZ5TTBvVEFvazg3aXl3SFRrbFRsSjM4In0.eyJqdGkiOiJmYjI5MTZjYS0yMWIyLTQ1NjMtODlmMS0wYWExMzFlNmZhZmUiLCJleHAiOjE1Njc1NjkwNTYsIm5iZiI6MCwiaWF0IjoxNTY3NTY4NzU2LCJpc3MiOiJodHRwczovL3Nzby5wb255dGVzdC5jb206ODQ0My9hdXRoL3JlYWxtcy9wb255dGVzdCIsImF1ZCI6ImlMYWItd2VieCIsInN1YiI6ImRkMWM0YzdlLTQ0ZmItNGU4OC05OWQ1LWNjZDA1N2M3M2JkMSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImlMYWItd2VieCIsIm5vbmNlIjoiOTViZDVmMTgtZWQ3NS00OWVkLTg2ZGYtNWE2YzY2NTMzNzliIiwiYXV0aF90aW1lIjoxNTY3NTYwMjU5LCJzZXNzaW9uX3N0YXRlIjoiMTE0NjlmNDYtZTU3Yi00NTY3LWI4NTUtMDQwODEzNDg5YjRkIiwiYWNyIjoiMCIsImFsbG93ZWQtb3JpZ2lucyI6WyJodHRwOi8vaWxhYi5wb255dGVzdC5jb20iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJyZWFsbS1tYW5hZ2VtZW50Ijp7InJvbGVzIjpbInZpZXctcmVhbG0iLCJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctZXZlbnRzIiwibWFuYWdlLXVzZXJzIiwicXVlcnktcmVhbG1zIiwidmlldy11c2VycyIsInZpZXctY2xpZW50cyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS1ncm91cHMiLCJxdWVyeS11c2VycyJdfSwiaWxhYi1yZXBvcnQiOnsicm9sZXMiOlsiUk9MRV9VU0VSIl19LCJpTGFiLXNlcnZlciI6eyJyb2xlcyI6WyJST0xFX0FQSSIsIlJPTEVfVVNFUiJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCIsInBvbnlfb3JnX2NvZGUiOiIxMDAxXzAxIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoi5bygIOa9h-a9hyIsInByZWZlcnJlZF91c2VybmFtZSI6InpoYW5neGlhb3hpYW8iLCJnaXZlbl9uYW1lIjoi5bygIiwibG9jYWxlIjoiemgtQ04iLCJmYW1pbHlfbmFtZSI6Iua9h-a9hyJ9.a5ltGuMo6rzV-PLSwTcKknU0df9f4PA_pMXKVMn63zJMx7QmTmpVWEBNaR5AmWBh8D562U06yhIcKt6ClmL5AlYW_kivzCAPdgjOTVc8NBcve7AhrJid9VXBKDCU1tw5JmTcFHsJ1CIbpR28qVTPY2yL4LT64HfUx9waf64IfQfSW9R5n76HI1gx7qb1kUDpLdjUuqE24BCYX-3kwAeg3w0e3EV0AiMA94JAVAcZs0q7lUB7QaAuUKtey1FhhOZz3WczBQYKco81FI-aQK6SU1u4vbyXCrZGdRKkOd_3NtlKXfOk_Hpkrfml-eQRozu_QAeQfS8z9qlPLgJk92fksw");
			},
			success: function (resp) {
				if (resp.code !== '0') {
					if(silence)return;
					mini.unmask(document.body);
					mini.showMessageBox({"title": "\u9519\u8bef\u002e","buttons": [],"iconCls": "mini-messagebox-error","message": "详细信息","width":"620px","height":"700px","html": `<textarea style="overflow:scroll;word-wrap:normal;resize:none;" rows="16" cols="96">${resp.message}</textarea>`,"showModal": true } );
					return;
				}
				setTimeout(function () {
					let msg = ok(resp)||'';
					if(msg.length>0){
						app.$Message.success({ content: msg, duration: 1 });
					}
				}, 1);
			},
			error: function (ex) {
				prompt(JSON.stringify(ex,null,4),"[ERROR]:"+route);
			},
			complete:function(){
				mini.unmask();
				if(done)done();
			}
		});
	},10)
}

DoPostBlob=function(expUrl, params,fileName,ext,done){
	ext=ext || '.xlsx';
	fileName = fileName || `${mini.get('currBiz').getValue()}@${new Date().format(`yyyy-MM-dd#hh_mm_ss`)}`;
	fileName = fileName + ext;
	xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		var a;
		if (xhttp.readyState === 4 && xhttp.status === 200) {
			a = document.createElement('a');
			a.href = window.URL.createObjectURL(xhttp.response);
			a.download = fileName;
			a.style.display = 'none';
			document.body.appendChild(a);
			a.click();
			if(done)done();
		}
	};
	xhttp.open("POST", expUrl);
	xhttp.setRequestHeader("Content-Type", "application/json");
	xhttp.setRequestHeader("Authorization", "Bearer " + getCookie('jwt'));
	xhttp.responseType = 'blob';
	xhttp.send(JSON.stringify(params));
}



MapKeyEditor=function(key,handle){
	app.editor.addCommand(key, function () {handle();});
}
SetEditorValue=function(val){
	app.editor.setValue(val);
}
ClearEditorValue=function(){
	app.editor.setValue("");
}
GetEditorValue=function(){
	return app.editor.getValue();
}
InsertEditorContent =function(text,editor){
	editor=editor||app.editor;
	let selection = editor.getSelection()
	let range = new monaco.Range(selection.startLineNumber, selection.startColumn, selection.endLineNumber, selection.endColumn)
	let id = { major: 1, minor: 1 }
	let op = {identifier: id, range: range, text: text, forceMoveMarkers: false}
	editor.executeEdits(app.root, [op])
	editor.focus()
}
FocusEditor=function(){
	app.editor.focus()
}



function getStorage(key){
	if(!window.localStorage){
		console.log('浏览器不支持 localStorage');
	}else{
		return localStorage.getItem(key);
	}
}
var _account = '';
function PersonalizedLayout(account,id) {
	_account = window.location.pathname.replace('/', '') + "_" + account;
//	$(".mini-datagrid").each(function(a, b) {
		var b={"id":id};// zhangxx 去除遍历 17:00 2019/8/30
		var grid = mini.get(b.id);
		var cols = grid.columns;
		var json = eval(getStorage(_account + "_" + b.id));
		var settings = null;
		if (json) {
			settings = JSON.parse(json);
			for (var j = 0; j < settings.length; j++) {
				var a = settings[j];
				for (var i = 0; i < cols.length; i++) {
					if (cols[i]["_id"] === a._id) {
						cols[i]._colIndex = j;
						break;
					}
				}
			}
			for (var i = 0; i < cols.length - 1; i++) {
				for (var j = 0; j < cols.length - 1 - i; j++) {
					if (cols[j]._colIndex > cols[j + 1]._colIndex) {
						var temp;
						temp = cols[j + 1];
						cols[j + 1] = cols[j];
						cols[j] = temp;
					}
				}
			}
		} //if(json)
		var menu = mini.create({
			type: "menu",
			hideOnClick: false
		});
		var items = [];
		for (var i = 0; i < cols.length; i++) {
			var setting = null;
			if (settings) if (settings[i]) setting = settings[i];
			var vis = true;
			if (setting) {
				grid.updateColumn(cols[i], {
					width: parseInt(setting.width)
				});
				vis = setting.visible !== false;
				if (vis) grid.showColumn(grid.getColumn(i));
				else grid.hideColumn(grid.getColumn(i));
			}
			var item = {
				"grid": b.id,
				"col": i
			};
			item.checked = vis;
			item.checkOnClick = true;
			item.text = '勾选'; //type:"checkcolumn"
			if ('string' === (typeof cols[i].header)) {
				if (cols[i].header.length > 0) {
					item.text = cols[i].header;
				}
			}
			items.push(item);
		}
		menu.setItems(items);
		menu.on("itemclick", "columnSetter", this);
		grid.setHeaderContextMenu(menu);
		grid.on("columnschanged", "gridConfiger");
//	});
}




function putStorage(key,value){
	if(!window.localStorage){
		console.log('浏览器不支持 localStorage');
	}else{
		localStorage[key]=JSON.stringify(value);
	}
}

function getStorage(key){
	if(!window.localStorage){
		console.log('浏览器不支持 localStorage');
	}else{
		return localStorage.getItem(key);
	}
}
// zhangxx 16:03 2018/12/4 表格加入动态表头菜单,实现自定义布局保存功能
function columnSetter(e) {
	var item = e.item;
	var grid=mini.get(item.grid);
	var col=grid.getColumn(item.col);
	if(!item.checked)
		grid.hideColumn(col);
	else
		grid.showColumn(col);
}

function gridConfiger(e){
	$(".mini-datagrid").each(function(a,b){
//if(b.id==='requestsGrid'){
		var grid=mini.get(b.id);
		var cols=grid.columns;
		
if(cols){

		var habits=[];
		for(var i=0;i<cols.length;i++){
			var col= grid.getColumn(i);
			//var col=cols[i];
			//if('string'!=(typeof col.header))continue;
/* // cookie 算法
			var key=b.id+"_"+i;
			document.cookie=key+"_width="+col.width;
			document.cookie=key+"_visible="+col.visible;
			document.cookie=key+"_colIndex="+col._id;
*/
			var habit={};
			habit["width"]=col.width;
			habit["visible"]=col.visible;
			habit["_id"]=col._id;
			habit["title"]="";
			if('string'===(typeof cols[i].header)){
				if(col.header.length>0){
					habit["title"]=col.header;
				}
			}

			//console.log(key,col._id,b.id,col.header,col.visible,col.width);
			habits.push(habit);
//}
		}
		putStorage(_account+"_"+b.id,JSON.stringify(habits,null,4));
}
	});
	if(new Date().getTime()-_gridConfiger>1500){
		//Toast.success('表格布局已保存');
		_gridConfiger=new Date().getTime();
	}
}
var _gridConfiger=0;

var _=null;
window_onload=function(){
	window.base64 = new Base64();
	window._=mini.get;
	$(function(){if(app){

		

		app.editor=monaco.editor.create(document.getElementById('codeContainer'), {language: 'sql',wrappingColumn: 0,wrappingIndent: "indent",scrollbar: {vertical: 'auto',horizontal: 'auto'},theme: "vs-dark",automaticLayout: true,readOnly: false,value: ''});
		app.editor.onDidChangeModelContent((e) => { });



		app.dsnRenderer=function(e){
			return app.Dsns[e.value];
		}
		app.dateRenderer=function(e){
			if(e.value)	return new Date((e.value||0).replace(/(\+\d{2})(\d{2})$/, "$1:$2")).format('yyyy-MM-dd hh:mm:ss');
		};
		app.iso8601Renderer=function(e){
			if(e.value)return new Date(e.value).format('yyyy-MM-dd hh:mm:ss');
		}
		mini.parse();
		hljs.initHighlightingOnLoad();
		if(app.init)app.init();


	}
	$(document.body).css("-webkit-transform","scale(1.0)");
	$(document).keydown(function (e) {
		if (27 === event.which)$('.mini-window').each(function(){if (undefined != mini.get($(this).attr('id')))mini.get($(this).attr('id')).hide();});
		else if(event.altKey){
			// if(65 == event.which)addSample();   （<u>A</u>）
		}
	});
	});
	document.onkeydown=function(){
		if(event.ctrlKey && event.keyCode == 82 || //Ctrl + R
			event.ctrlKey && event.keyCode == 83) { //Ctrl + S
			event.keyCode = 0;
			event.returnValue = false;
			return;
		}
	}
}









function getCookie(sName) {
	var aCookie = document.cookie.split("; ");
	var lastMatch = null;
	for (var i = 0; i < aCookie.length; i++) {
		var aCrumb = aCookie[i].split("=");
		if (sName == aCrumb[0]) {
			lastMatch = aCrumb;
		}
	}
	if (lastMatch) {
		var v = lastMatch[1];
		if (v === undefined) return v;
		return unescape(v);
	}
	return null;
}
/////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////
var debug=false;
//zhangxx 9:18 2018/10/24
(function () {
	var ie = !!(window.attachEvent && !window.opera);
	var wk = /webkit\/(\d+)/i.test(navigator.userAgent) && (RegExp.$1 < 525);
	var fn = [];
	var run = function () { for (var i = 0; i < fn.length; i++) fn[i](); };
	var d = document;
	d.ready = function (f) {
	if (!ie && !wk && d.addEventListener)
	return d.addEventListener('DOMContentLoaded', f, false);
	if (fn.push(f) > 1) return;
	if (ie)
		(function () {
			try { d.documentElement.doScroll('left'); run(); }
			catch (err) { setTimeout(arguments.callee, 0); }
		})();
	else if (wk)
	var t = setInterval(function () {
		if (/^(loaded|complete)$/.test(d.readyState))
		clearInterval(t), run();
	}, 0);
	};
})();
//if(undefined===window.keycloak){
function jwtSet(){
	if (typeof jQuery == 'undefined') {
		setTimeout(jwtSet,10);
		return;
	}
	$.ajaxSetup({
		beforeSend : function(xhr) {
			xhr.setRequestHeader("Authorization", "Bearer "+ getCookie('jwt'));
		},
	});
}


if("localhost"===window.location.hostname){
	var exp = new Date();
	exp.setTime(exp.getTime() + 120 * 60 * 1000);
	let mockToken = 'eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI3SVJUX3N6bkVkc3hIZUl0V3NCZmZ5TTBvVEFvazg3aXl3SFRrbFRsSjM4In0.eyJqdGkiOiJjNzAyYTM1NC0yMmE4LTRlM2UtOGM4ZS01M2JkZTM2NzU1MDQiLCJleHAiOjE1Njg3MDIwMjIsIm5iZiI6MCwiaWF0IjoxNTY4NzAxNzIyLCJpc3MiOiJodHRwczovL3Nzby5wb255dGVzdC5jb206ODQ0My9hdXRoL3JlYWxtcy9wb255dGVzdCIsImF1ZCI6ImlMYWItd2VieCIsInN1YiI6ImRkMWM0YzdlLTQ0ZmItNGU4OC05OWQ1LWNjZDA1N2M3M2JkMSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImlMYWItd2VieCIsIm5vbmNlIjoiOTNmOTYyZDUtZDg4NS00YjdiLTg2MGUtOWYzNTMwODhkN2I1IiwiYXV0aF90aW1lIjoxNTY4NjgxODAyLCJzZXNzaW9uX3N0YXRlIjoiYjNlMjNhMGYtOTk5Ni00N2ViLWFhOWEtZmRiYTM3ZDFjNzQ4IiwiYWNyIjoiMCIsImFsbG93ZWQtb3JpZ2lucyI6WyJodHRwOi8vaWxhYi5wb255dGVzdC5jb20iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJyZWFsbS1tYW5hZ2VtZW50Ijp7InJvbGVzIjpbInZpZXctcmVhbG0iLCJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctZXZlbnRzIiwibWFuYWdlLXVzZXJzIiwicXVlcnktcmVhbG1zIiwidmlldy11c2VycyIsInZpZXctY2xpZW50cyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS1ncm91cHMiLCJxdWVyeS11c2VycyJdfSwiaWxhYi1yZXBvcnQiOnsicm9sZXMiOlsiUk9MRV9VU0VSIl19LCJpTGFiLXNlcnZlciI6eyJyb2xlcyI6WyJST0xFX0FQSSIsIlJPTEVfVVNFUiJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19LCJpbGFiLXJlcG9ydC14Ijp7InJvbGVzIjpbIlJPTEVfVVNFUiJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJwb255X29yZ19jb2RlIjoiMTAwMV8wMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6IuW8oCDmvYfmvYciLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ6aGFuZ3hpYW94aWFvIiwiZ2l2ZW5fbmFtZSI6IuW8oCIsImxvY2FsZSI6InpoLUNOIiwiZmFtaWx5X25hbWUiOiLmvYfmvYcifQ.eh3cB-3vRErquBp1v8eQbw2SJqJccNS6taVpW1jQ6QgvjskSNK3DNNuqCUFwSKjlF1kpy7DIB7LASrnVcp4JQ6bT3RyPKqNF8PxrCLxEHHalsNZx7ZDANCBbs8bd2Q-O7H6m0kC3lwgR1tYt8EcTXN0JzXRq9EX0kYsPe09lv28MwKO-BnxpgUtHxB0FStf5Uijd06EIQdurt2I6k0hdirRIfd-5kxU8t9wfhs9ZHlbeXuZyzVFd-CH_-LKWcXtdZwmzPzBEeWcwqqa7WMENpQMjnpueq4d_JHN2TqdlYeh7Q6_0uSvtuRCJghAtJOlAoxYPN5Wi8O4KpmFNRvn-lg';
	document.cookie=`jwt=${mockToken};expires=${exp.toGMTString()}`;
	jwtSet();
	setTimeout(window_onload,800);
}else{
	jwtSet();
	//console.log('window.location.pathname',window.location.pathname);
	if(null==getCookie('jwt')||'/reportx/'===window.location.pathname){
		document.write('<script src="' + 'https://sso.ponytest.com:8443/auth/js/keycloak.min.js' + '?rnd=' + Math.random() + '" type="text/javascript"></sc' + 'ript>');
		document.ready(function(){
			function output(data) {
				if (typeof data === 'object') data = JSON.stringify(data, null, '  ');
				if(debug)console.log('oauth:'+data);
			}
			if(undefined!==window.keycloak)return;
			window.keycloak = Keycloak();
			window.keycloak.onAuthSuccess = function () {output('登录成功');};
			window.keycloak.onAuthError = function (errorData) {output("登录失败：" + JSON.stringify(errorData) );};
			window.keycloak.onAuthRefreshSuccess = function () {output("token 刷新成功");};
			window.keycloak.onAuthRefreshError = function () {output('token 刷新失败');};
			window.keycloak.onAuthLogout = function () {output('注销成功');};
			window.keycloak.onTokenExpired = function () {output('token过期');};
			window.keycloak.init({
				responseMode: 'fragment', //可选值：fragment、query
				flow: 'standard',//可选值：standard、implicit、hybrid
				onLoad: 'check-sso' //可选值：check-sso、login-required、或不配置
			}).success(function(authenticated) {
				output('初始化：' + (authenticated ? 'Authenticated' : 'Not Authenticated'));
				function setter(){
					output(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"+window.keycloak.token);
					console.log('setter()'+new Date().toLocaleString());
					var exp = new Date();
					exp.setTime(exp.getTime() + 120 * 60 * 1000);
					document.cookie='jwt='+window.keycloak.token + ";expires=" + exp.toGMTString()
					jwtSet();
				}
				if(authenticated){
					setter();
					window_onload();
					// setTimeout(function () {
					//     //window.location.href = 'desktop.html?accessKey=aabbccdd';
					// }, 1200);
					setInterval(function(){
						window.keycloak.updateToken(60).success(function(refreshed) {
							if (refreshed) {
								//output(window.keycloak.tokenParsed);
								setter();
							} else {
								output('Token not refreshed, valid for ' + Math.round(window.keycloak.tokenParsed.exp + window.keycloak.timeSkew - new Date().getTime() / 1000) + ' seconds');
							}
						}).error(function() {
							output('Failed to refresh token');
						});
					},60*1000);
				}else{
					window.keycloak.login();
				}
			}).error(function() {
				console.log('验证身份错误');
			});
		});//documentReady
	}
}