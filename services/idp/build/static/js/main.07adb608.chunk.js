(this.webpackJsonpidentifier=this.webpackJsonpidentifier||[]).push([[6],{43:function(e){e.exports=JSON.parse('[{"locale":"de","name":"German","nativeName":"Deutsch","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"en-GB","name":"English","nativeName":"English","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"es","name":"Spanish","nativeName":"Espa\xf1ol","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"fr","name":"French","nativeName":"Fran\xe7ais","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"it","name":"Italian","nativeName":"Italiano","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"nl","name":"Dutch","nativeName":"Nederlands","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"}},{"locale":"zh-CN","name":"Simplified Chinese","nativeName":"\u7b80\u4f53\u4e2d\u6587","orientation":{"characterOrder":"left-to-right","lineOrder":"top-to-bottom"},"ietf":"zh_hans"}]')},5:function(e,n,t){"use strict";t.d(n,"b",(function(){return r})),t.d(n,"k",(function(){return o})),t.d(n,"c",(function(){return a})),t.d(n,"f",(function(){return i})),t.d(n,"j",(function(){return c})),t.d(n,"e",(function(){return l})),t.d(n,"m",(function(){return s})),t.d(n,"g",(function(){return u})),t.d(n,"h",(function(){return d})),t.d(n,"a",(function(){return O})),t.d(n,"i",(function(){return f})),t.d(n,"d",(function(){return E})),t.d(n,"l",(function(){return h}));var r="RECEIVE_ERROR",o="RESET_HELLO",a="RECEIVE_HELLO",i="RECEIVE_VALIDATE_LOGON",c="REQUEST_LOGON",l="RECEIVE_LOGON",s="UPDATE_INPUT",u="REQUEST_CONSENT_ALLOW",d="REQUEST_CONSENT_CANCEL",O="RECEIVE_CONSENT",f="REQUEST_LOGOFF",E="RECEIVE_LOGOFF",h="SERVICE_WORKER_NEW_CONTENT"},90:function(e,n,t){var r={"./de/translation.json":104,"./es/translation.json":105,"./fr/translation.json":106,"./it/translation.json":107,"./locales.json":43,"./nl/translation.json":108,"./zh-CN/translation.json":109};function o(e){return a(e).then((function(e){return t.t(e,3)}))}function a(e){return t.e(1).then((function(){if(!t.o(r,e)){var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}return r[e]}))}o.keys=function(){return Object.keys(r)},o.resolve=a,o.id=90,e.exports=o},93:function(e,n,t){},97:function(e,n,t){},98:function(e,n,t){"use strict";t.r(n);var r=t(0),o=t.n(r),a=t(13),i=t.n(a),c=t(49),l=t(15),s=t(46),u=t(103),d=t(56),O=t.n(d),f=t(57),E=t(29),h=t.n(E),b=t(43);var m="ui_locales",p="ui_locale",_="lico.identifier_ui_locale",g=b.map((function(e){return e.locale})),j={name:"queryUiLocales",lookup:function(e){var n=h.a.parse(document.location.search)[m];if(n)return Array.isArray(n)?n:n.split(" ")}},v=new f.a;v.addDetector(j),s.a.use(O()((function(e,n,r){t(90)("./".concat(e,"/").concat(n,".json")).then((function(e){r(null,e)})).catch((function(e){r(e,null)}))}))).use(v).use(u.e).init({fallbackLng:"en-GB",supportedLngs:Object(l.a)(g),cleanCode:!0,returnEmptyString:!1,debug:!1,detection:{order:[j.name,"cookie","localStorage","navigator"],lookupCookie:p,lookupLocalStorage:_,caches:["localStorage"]},interpolation:{escapeValue:!1}});s.a;var N=t(125),S=t(66),C=(t(91),t(92),t(93),t(123)),T=t(124),R=t(121),L=t(6),y=Object(R.a)((function(){return{root:{position:"fixed",top:"50%",left:"50%",transform:"translate(-50%, -50%)"}}})),I=function(){var e=y();return Object(L.jsx)("div",{className:e.root,children:Object(L.jsx)(C.a,{in:!0,style:{transitionDelay:"800ms"},unmountOnExit:!0,children:Object(L.jsx)(T.a,{size:70,thickness:1})})})},x=Object({NODE_ENV:"production",PUBLIC_URL:".",WDS_SOCKET_HOST:void 0,WDS_SOCKET_PATH:void 0,WDS_SOCKET_PORT:void 0,FAST_REFRESH:!0}).REACT_APP_KOPANO_BUILD||"0.0.0-no-proper-build",A=Object(r.lazy)((function(){return Promise.all([t.e(9),t.e(5)]).then(t.bind(null,232))}));console.info("Kopano Identifier build version: ".concat(x));var D=function(){return Object(L.jsx)(N.a,{theme:S.a,children:Object(L.jsx)(r.Suspense,{fallback:Object(L.jsx)(I,{}),children:Object(L.jsx)(A,{})})})},k=t(19),w=t(64),P=(t(96),t(5)),U=h.a.parse(document.location.search),V=U.flow||"";delete U.flow;var F={hello:null,branding:null,error:null,flow:V,query:U,updateAvailable:!1,pathPrefix:function(){var e=document.getElementById("root"),n=e?e.getAttribute("data-path-prefix"):null;return n&&"__PATH_PREFIX__"!==n||(n="/signin/v1"),n}()};var G=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:F,n=arguments.length>1?arguments[1]:void 0;switch(n.type){case P.b:return Object.assign({},e,{error:n.error});case P.k:return Object.assign({},e,{hello:null,branding:null});case P.c:return Object.assign({},e,{hello:{state:n.state,username:n.username,displayName:n.displayName,details:n.hello},branding:n.hello.branding?n.hello.branding:e.branding});case P.l:return Object.assign({},e,{updateAvailable:!0});default:return e}},B=t(4);var H=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{loading:"",username:"",password:"",errors:{}},n=arguments.length>1?arguments[1]:void 0;switch(n.type){case P.f:return Object.assign({},e,{errors:n.errors,loading:""});case P.g:case P.h:case P.j:return Object.assign({},e,{loading:n.type,errors:{}});case P.a:case P.e:return n.success?e:Object.assign({},e,{errors:n.errors?n.errors:{},loading:""});case P.d:return Object.assign({},e,{username:"",password:""});case P.m:return delete e.errors[n.name],Object.assign({},e,Object(B.a)({},n.name,n.value));default:return e}},K=Object(k.b)({common:G,login:H}),W=[w.a];var z=window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__||k.c,Q=Object(k.d)(K,z(k.a.apply(void 0,W)));t(97);i.a.render(Object(L.jsx)(o.a.StrictMode,{children:Object(L.jsx)(c.a,{store:Q,children:Object(L.jsx)(D,{})})}),document.getElementById("root"))}},[[98,7,8]],[1]]);
//# sourceMappingURL=main.07adb608.chunk.js.map