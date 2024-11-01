"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[2634],{7735:(e,i,t)=>{t.r(i),t.d(i,{default:()=>Y});var o=t(6540),n=t(8206),r=t(8774),s=t(4586),a=t(3705),d=t(4012),c=t(6844),l=t(4798),m=t(1118),g=t(4848);const p=a.A.div`
  display: flex;
  flex-direction: column;
  max-width: 1200px;
  margin: 4rem auto;
  padding: 0 1rem;
`,h=a.A.div`
  display: flex;
  flex-direction: row;
  gap: 48px;

  @media (max-width: 768px) {
    flex-direction: column;
  }
`,u=a.A.div`
  flex: 1;
`,x=a.A.h2`
  font-size: 1.5rem;
  margin-bottom: 1rem;
`,f=a.A.p`
  font-size: 1rem;
  color: var(--ifm-color-emphasis-700);
  margin-bottom: 2rem;
`,k=(0,a.A)(r.A)`
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    margin-bottom: 0.5rem;
    border-radius: 8px;
    background-color: var(--ifm-background-surface-color);
    text-decoration: none;
    color: var(--ifm-font-color-base);
    transition: background-color 0.2s;

    &:hover {
        background-color: var(--ifm-color-emphasis-100);
        text-decoration: none; /* Prevents underline on hover */
    }
`,b=a.A.h3`
  font-size: 1rem;
  margin: 0;
`,w=a.A.p`
  font-size: 0.875rem;
  color: var(--ifm-color-emphasis-600);
  margin: 0.25rem 0 0;
`,j=[{title:"Get started",description:"Step-by-step guide to install and configure the necessary tools",to:"/docs/quickstart/install-guide"},{title:"Complete Guide",description:"Follow a comprehensive guide to set up your environment",to:"/docs/quickstart/complete-guide"},{title:"Check Dependencies",description:"Check the dependencies required for your setup",to:"/docs/quickstart/dependencies"},{title:"Keys Management",description:"Manage your validator keys efficiently",to:"/docs/quickstart/keys-management"},{title:"Commands",description:"Learn about advanced commands to manage or set up your node",to:"/docs/commands"}],v=[{title:"Ethereum Mainnet",description:"Deploy and manage your nodes on Ethereum Mainnet",to:"/docs/networks/mainnet"},{title:"Gnosis Network",description:"Get started with the Gnosis Mainnet and Testnet",to:"/docs/networks/gnosis"},{title:"Optimism Network",description:"Deploy on Optimism Mainnet and its test networks",to:"/docs/quickstart/running-optimism-node"},{title:"Exposing API",description:"Guide on exposing APIs for your validator setup",to:"/docs/quickstart/samples/exposing-apis"},{title:"Using Relays",description:"Learn how to utilize relays in your setup",to:"/docs/quickstart/samples/using-relays"}];function y(){return(0,g.jsx)(p,{children:(0,g.jsxs)(h,{children:[(0,g.jsxs)(u,{children:[(0,g.jsx)(x,{children:"Integrate with Sedge"}),(0,g.jsx)(f,{children:"Explore these guided tutorials to get started with Sedge for your Ethereum staking needs."}),j.map((e=>(0,g.jsxs)(k,{to:e.to,children:[(0,g.jsxs)("div",{children:[(0,g.jsx)(b,{children:e.title}),(0,g.jsx)(w,{children:e.description})]}),(0,g.jsx)(m.A,{size:20})]},e.title)))]}),(0,g.jsxs)(u,{children:[(0,g.jsx)(x,{children:"Advanced Sedge Features"}),(0,g.jsx)(f,{children:"Dive deeper into Sedge's capabilities and optimize your Ethereum staking experience."}),v.map((e=>(0,g.jsxs)(k,{to:e.to,children:[(0,g.jsxs)("div",{children:[(0,g.jsx)(b,{children:e.title}),(0,g.jsx)(w,{children:e.description})]}),(0,g.jsx)(m.A,{size:20})]},e.title)))]})]})})}var A=t(1122),S=t(6025);const E=a.A.h2`
    text-align: left;
    margin-left: 1rem;
    margin-bottom: 2rem;
    font-size: 1.5rem;
`,z=a.A.div`
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 1rem;
`,q=a.A.div`
    display: flex;
    flex-direction: column;
    padding-right: 4rem;
    padding-left: 4rem;
    padding-top: 2rem;
    padding-bottom: 1rem;
    border: 1px solid var(--ifm-color-emphasis-300);
    border-radius: 8px;
    text-align: center;
    background-color: var(--ifm-background-surface-color);
`,N=(0,a.A)(A.A)`
    width: 80px;
    height: 80px;
    object-fit: contain;
    margin: 0 auto 0.5rem;
`,C=a.A.h3`
    margin-bottom: 0.5rem;
    font-size: 1.1rem;
    color: var(--ifm-color-emphasis-900);
`,L=a.A.a`
    display: inline-block;
    margin: 0.25rem;
    padding: 0.5rem 1rem;
    background-color: var(--ifm-color-emphasis-200);
    color: var(--ifm-color-emphasis-700);
    border-radius: 4px;
    text-decoration: none;
    font-size: 0.9rem;
`,M=a.A.div`
    display: flex;
    flex-direction: column;
    align-items: stretch;
    gap: 0.5rem;
`,D=(0,a.A)(L)`
    display: block;
    margin: 0;
`,G=()=>(0,g.jsxs)(g.Fragment,{children:[(0,g.jsx)(E,{children:"Supported Networks"}),(0,g.jsx)(z,{children:[{name:"Base",logo:{light:"/img/chains/base-logo.png",dark:"/img/chains/base-logo.png"},networks:[{name:"Mainnet",link:"/docs/quickstart/running-optimism-node#base-support"}]},{name:"Ethereum",logo:{light:"/img/chains/eth-logo.svg",dark:"/img/chains/eth-logo.svg"},networks:[{name:"Mainnet",link:"/docs/networks/mainnet"}]},{name:"Gnosis",logo:{light:"/img/chains/gno-logo.png",dark:"/img/chains/gno-logo.png"},networks:[{name:"Mainnet",link:"/docs/networks/gnosis"}]},{name:"Optimism",logo:{light:"/img/chains/op-logo.png",dark:"/img/chains/op-logo.png"},networks:[{name:"Mainnet",link:"/docs/quickstart/running-optimism-node"}]},{name:"Testnets",networks:[{name:"Holesky",link:"/docs/networks/holesky"},{name:"Sepolia",link:"/docs/networks/sepolia"},{name:"Chiado",link:"/docs/networks/chiado"}]},{name:"Lido",logo:{light:"/img/chains/lido-logo.png",dark:"/img/chains/lido-logo.png"},networks:[{name:"CSM",link:"/docs/quickstart/staking-with-lido"}]}].map((e=>(0,g.jsxs)(q,{children:[e.logo&&(0,g.jsx)(N,{sources:{light:(0,S.Ay)(e.logo.light),dark:(0,S.Ay)(e.logo.dark)},alt:`${e.name} logo`}),(0,g.jsx)(C,{children:e.name}),"Testnets"===e.name?(0,g.jsx)(M,{children:e.networks.map((e=>(0,g.jsx)(D,{href:e.link,children:e.name},e.name)))}):e.networks.map((e=>(0,g.jsx)(L,{href:e.link,children:e.name},e.name)))]},e.name)))})]});var P=t(3546),T=t(5773),F=t(5404);const H=e=>{let{command:i}=e;const[t,n]=(0,o.useState)(!1);return(0,g.jsx)("div",{className:"bash-command-container",children:(0,g.jsxs)("div",{className:"bash-command-content",children:[(0,g.jsx)("div",{className:"bash-command-code",children:(0,g.jsx)("pre",{children:(0,g.jsx)("code",{children:i})})}),(0,g.jsx)("div",{children:(0,g.jsx)("button",{onClick:async()=>{try{await navigator.clipboard.writeText(i),n(!0),setTimeout((()=>n(!1)),2e3)}catch(e){console.error("Failed to copy text: ",e)}},className:"bash-command-button","aria-label":"Copy to clipboard",children:t?(0,g.jsx)(T.A,{className:"bash-command-icon",style:{color:"green"}}):(0,g.jsx)(F.A,{className:"bash-command-icon",style:{color:"gray"}})})})]})})},I=()=>{const[e,i]=o.useState("linux-macos");o.useEffect((()=>{const e=navigator.userAgent||navigator.vendor||window.opera;/windows phone/i.test(e)||/win/i.test(e)?i("windows"):(/android/i.test(e)||/iPad|iPhone|iPod/.test(e)||/mac/i.test(e)||/linux/i.test(e))&&i("linux-macos")}),[]);return(0,g.jsxs)("div",{children:[(0,g.jsxs)("p",{children:["To install ",(0,g.jsx)("strong",{children:"Sedge"}),", run the following command in your terminal:"]}),(0,g.jsx)(H,{command:{"linux-macos":"bash <(curl -fsSL https://github.com/NethermindEth/sedge/raw/main/scripts/install.sh)",windows:"Set-ExecutionPolicy Bypass -Scope Process -Force; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://github.com/NethermindEth/sedge/raw/main/scripts/install.ps1'))"}[e]})]})};function O(){const e=a.A.div`
        width: 100%;
        //background-image: url('/img/background-image.jpg');
        background-size: cover;
        background-position: center;
        position: relative;

        //&::before {
        //    content: '';
        //    position: absolute;
        //    top: 0;
        //    left: 0;
        //    right: 0;
        //    bottom: 0;
        //    background: rgba(0, 0, 0, 0.7); // Darker overall overlay
        //}

    `,i=a.A.header`
        position: relative;
        padding: 4rem 0;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
        max-width: var(--ifm-container-width);
        margin: 0 auto;
        color: var(--ifm-font-color-base);
    `,t=a.A.h1`
        font-weight: 600;
        font-size: 3rem;
        margin-top: 1rem;
        margin-bottom: 1rem;
    `,o=a.A.div`
        margin-top: 2rem;
        width: 100%;
        max-width: 600px;
        display: flex;
        justify-content: center;
    `;return(0,g.jsx)(e,{children:(0,g.jsxs)(i,{children:[(0,g.jsx)(A.A,{sources:{dark:(0,S.Ay)("/img/Sedge_Horizontal_Light.svg"),light:(0,S.Ay)("/img/Sedge_Horizontal_Dark.svg")},alt:"Sedge Logo"}),(0,g.jsx)(t,{children:"Easy node setup and deployment tool"}),(0,g.jsx)(I,{}),(0,g.jsx)(o,{children:(0,g.jsx)(P.A,{})})]})})}const _=a.A.div`
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    min-height: 100vh;
`,$=a.A.div`
    position: fixed;
    top: 0;
    left: 0;
    height: 4px;
    background-color: var(--ifm-color-primary-dark);
    z-index: 1000;
`,B=[{title:"What is Sedge",icon:d.A,to:"/docs/intro",text:"Learn about the core concepts of Sedge, its features, and how it can help you."},{title:"Get Started with Sedge",icon:c.A,to:"/docs/quickstart",text:"Learn how to install and set up Sedge for your Ethereum staking needs."},{title:"Sedge Documentation",icon:l.A,to:"/docs/quickstart/install-guide",text:"Explore the full documentation to learn about all Sedge features and capabilities."}];function W(){const e=a.A.div`
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    grid-gap: 24px;
    justify-content: center;
    margin: 0 auto;
    padding: 2rem;
    max-width: 1200px;
    //background-color: var(--ifm-background-surface-color);
  `,i=a.A.div`
    display: flex;
    padding: 1.5rem;
    flex-direction: column;
    justify-content: space-between;
    cursor: pointer;
    border: 1px solid var(--ifm-color-emphasis-300);
    border-radius: 12px;
    transition: all 0.3s;
    background-color: var(--ifm-background-color);
    margin-bottom: 1rem;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
    }
  `,t=a.A.h3`
    margin-bottom: 1rem;
    font-weight: 600;
  `,o=a.A.p`
    margin-bottom: 1rem;
    font-weight: 400;
  `,n=a.A.div`
    width: 48px;
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    background-color: var(--ifm-color-emphasis-100);
    margin-bottom: 1rem;
  `;return(0,g.jsx)(e,{children:B.map((e=>(0,g.jsxs)(i,{children:[(0,g.jsxs)("div",{children:[(0,g.jsx)(n,{children:(0,g.jsx)(e.icon,{size:24})}),(0,g.jsx)(t,{children:e.title}),(0,g.jsx)(o,{children:e.text})]}),(0,g.jsxs)(r.A,{to:e.to,style:{display:"flex",alignItems:"center",color:"var(--ifm-color-primary)"},children:["Learn more ",(0,g.jsx)(m.A,{size:16,style:{marginLeft:"0.5rem"}})]})]},e.title)))})}function Y(){const{siteConfig:e}=(0,s.A)(),[i,t]=(0,o.useState)(0);return(0,o.useEffect)((()=>{const e=()=>{const e=document.documentElement.scrollHeight-window.innerHeight,i=window.scrollY/e*100;t(i)};return window.addEventListener("scroll",e),()=>window.removeEventListener("scroll",e)}),[]),(0,g.jsxs)(n.A,{title:`${e.title} Documentation`,description:"Technical Documentation For Sedge",children:[(0,g.jsx)($,{style:{width:`${i}%`}}),(0,g.jsx)(O,{}),(0,g.jsxs)(_,{children:[(0,g.jsx)(W,{}),(0,g.jsx)(G,{}),(0,g.jsx)(y,{})]})]})}}}]);