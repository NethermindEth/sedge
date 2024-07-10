"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[2253],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var l=a.createContext({}),u=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},p=function(e){var t=u(e.components);return a.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},g=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),g=u(n),m=r,c=g["".concat(l,".").concat(m)]||g[m]||d[m]||o;return n?a.createElement(c,i(i({ref:t},p),{},{components:n})):a.createElement(c,i({ref:t},p))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=g;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:r,i[1]=s;for(var u=2;u<o;u++)i[u]=n[u];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}g.displayName="MDXCreateElement"},7977:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>d,frontMatter:()=>o,metadata:()=>s,toc:()=>u});var a=n(7462),r=(n(7294),n(3905));const o={id:"migrate-setup",sidebar_position:11},i="Migrate from another setup",s={unversionedId:"quickstart/samples/migrate-setup",id:"quickstart/samples/migrate-setup",title:"Migrate from another setup",description:"If you already have a setup running and want to migrate it to a new server, change the client you are using, or upgrade",source:"@site/docs/quickstart/samples/migrate-current-setup.mdx",sourceDirName:"quickstart/samples",slug:"/quickstart/samples/migrate-setup",permalink:"/docs/quickstart/samples/migrate-setup",draft:!1,editUrl:"https://github.com/NethermindEth/sedge/tree/main/docs/docs/quickstart/samples/migrate-current-setup.mdx",tags:[],version:"current",sidebarPosition:11,frontMatter:{id:"migrate-setup",sidebar_position:11},sidebar:"tutorialSidebar",previous:{title:"Running a MEV-boost node",permalink:"/docs/quickstart/samples/running-mev-boost-node"},next:{title:"Staking with Lido using Sedge",permalink:"/docs/quickstart/staking-with-lido"}},l={},u=[{value:"Migrating from a not Sedge setup",id:"migrating-from-a-not-sedge-setup",level:2},{value:"Migrating from a Sedge setup to another Sedge setup",id:"migrating-from-a-sedge-setup-to-another-sedge-setup",level:2}],p={toc:u};function d(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"migrate-from-another-setup"},"Migrate from another setup"),(0,r.kt)("p",null,"If you already have a setup running and want to migrate it to a new server, change the client you are using, or upgrade\nyour node, you are in the right place."),(0,r.kt)("admonition",{type:"tip"},(0,r.kt)("p",{parentName:"admonition"},"We recommend you to investigate and know which clients you are going to use before migrating your setup. This will help\nyou to avoid some issues that may arise during the migration process.")),(0,r.kt)("h2",{id:"migrating-from-a-not-sedge-setup"},"Migrating from a not Sedge setup"),(0,r.kt)("p",null,"If you are migrating from a setup that is not Sedge, you will need to do the following:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Generate a new setup using Sedge, you can choose between the interactive or the non-interactive setup.\nIn this case we will use the non-interactive setup.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge generate full-node\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Generate the keys, if you don't have ones already.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge keys\n")),(0,r.kt)("p",null,"If you already have keys, you can copy them to the new setup. The keys are usually located in the\n",(0,r.kt)("inlineCode",{parentName:"p"},"./sedge-data/keystore")," folder. It might require that the keys are in the expected format."),(0,r.kt)("p",null,"Otherwise, you can use your mnemonic to generate the keys, using the same ",(0,r.kt)("inlineCode",{parentName:"p"},"sedge keys")," command."),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Import keys, either by copying them to the new setup or using the one generated in the previous step. The format of\nthe keys should follow the EIP-2335 standard.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge import-key\n")),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"If you are using any politics for slashing protection on your node, you can export that info following the slashing export\ninstructions of the node you are using. It will need to follow the EIP-3076 format.")),(0,r.kt)("p",null,"Once you have the JSON file, you can import it to the new setup."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge slashing-import --from slashing-data.json [validator]\n")),(0,r.kt)("p",null,"You will need to provide the validator client you are going to use in your setup, in order to import the slashing data."),(0,r.kt)("ol",{start:5},(0,r.kt)("li",{parentName:"ol"},"If you follow the previous steps, you should have a new setup with the same keys and slashing protection data.")),(0,r.kt)("p",null,"You are ready to run it:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge run\n")),(0,r.kt)("h2",{id:"migrating-from-a-sedge-setup-to-another-sedge-setup"},"Migrating from a Sedge setup to another Sedge setup"),(0,r.kt)("p",null,"If you are migrating the setup from a previously generated Sedge setup, it will be easier."),(0,r.kt)("p",null,"You will be able to directly export the keys and slashing protection data, and import them to the new setup, using sedge\nwithout the need to follow any other step."),(0,r.kt)("p",null,"You can follow the next steps to make the migration:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Generate the new setup, you can choose between the interactive or the non-interactive setup.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge generate full-node --path new-path\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"You can either import the keystore folder from your previous setup, or generate the keys again using the mnemonic.\nIn this case we will use the ",(0,r.kt)("inlineCode",{parentName:"li"},"sedge import-key")," command.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"sedge import-key --from [old-keystore-dir]\n")),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Export the slashing protection data from the previous setup.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge slashing-export --out slashing-data.json [validator]\n")),(0,r.kt)("p",null,"You will need to provide the validator client you are using in your previous setup, in order to export the slashing data."),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"Import the slashing protection data to the new setup.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge slashing-import --from slashing-data.json [validator]\n")),(0,r.kt)("p",null,"You will need to provide the validator client you are going to use in your setup, in order to import the slashing data."),(0,r.kt)("ol",{start:5},(0,r.kt)("li",{parentName:"ol"},"If you follow the previous steps, you should have a new setup with the same keys and slashing protection data.")),(0,r.kt)("p",null,"You are ready to run it:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sedge run\n")))}d.isMDXComponent=!0}}]);