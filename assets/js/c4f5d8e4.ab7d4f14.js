"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[2634],{7735:(e,i,t)=>{t.r(i),t.d(i,{default:()=>B});var r=t(6540),a=t(8206),o=t(8774),n=t(4586),m=t(3705),d=t(4012),s=t(6844),p=t(4798),l=t(1118),x=t(4848);const c=m.A.div`
  display: flex;
  flex-direction: column;
  max-width: 1200px;
  width: 100%;
  margin: 4rem auto;
  padding: 0 1.5rem;

  @media (max-width: 1024px) {
    max-width: 900px;
    margin: 3rem auto;
  }

  @media (max-width: 768px) {
    margin: 2rem auto;
    padding: 0 1rem;
  }

  @media (max-width: 480px) {
    margin: 1.5rem auto;
    padding: 0 0.75rem;
  }
`,h=m.A.div`
  display: flex;
  flex-direction: row;
  gap: 48px;
  width: 100%;

  @media (max-width: 1024px) {
    gap: 32px;
  }

  @media (max-width: 768px) {
    flex-direction: column;
    gap: 40px;
  }

  @media (max-width: 480px) {
    gap: 32px;
  }
`,g=m.A.div`
  flex: 1;
  min-width: 0; // Prevents flex items from overflowing

  @media (max-width: 768px) {
    width: 100%;
  }

  & + & {
    @media (max-width: 768px) {
      padding-top: 1rem;
      border-top: 1px solid var(--ifm-color-emphasis-200);
    }
  }
`,w=m.A.h2`
  font-size: 1.5rem;
  margin-bottom: 1rem;

  @media (max-width: 768px) {
    font-size: 1.25rem;
    margin-bottom: 0.75rem;
  }
`,u=m.A.p`
  font-size: 1rem;
  color: var(--ifm-color-emphasis-700);
  margin-bottom: 2rem;

  @media (max-width: 768px) {
    font-size: 0.875rem;
    margin-bottom: 1.5rem;
  }
`,f=(0,m.A)(o.A)`
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1.25rem;
    margin-bottom: 0.75rem;
    border-radius: 8px;
    background-color: var(--ifm-background-surface-color);
    text-decoration: none;
    color: var(--ifm-font-color-base);
    transition: background-color 0.2s, transform 0.2s;
    width: 100%;

    @media (max-width: 1024px) {
      padding: 1.125rem;
    }

    @media (max-width: 768px) {
      padding: 1rem;
      margin-bottom: 0.5rem;
    }

    @media (max-width: 480px) {
      padding: 0.875rem;
      margin-bottom: 0.375rem;
    }

    @media (hover: hover) {
      &:hover {
        transform: translateY(-2px);
      }
    }

    &:active {
      transform: translateY(1px);
    }

    svg {
      flex-shrink: 0;
      margin-left: 1rem;

      @media (max-width: 480px) {
        width: 18px;
        height: 18px;
        margin-left: 0.75rem;
      }
    }

    &:hover {
        background-color: var(--ifm-color-emphasis-100);
        text-decoration: none; /* Prevents underline on hover */
    }
`,b=m.A.h3`
  font-size: 1rem;
  margin: 0;

  @media (max-width: 768px) {
    font-size: 0.9375rem;
  }
`,v=m.A.p`
  font-size: 0.875rem;
  color: var(--ifm-color-emphasis-600);
  margin: 0.25rem 0 0;

  @media (max-width: 768px) {
    font-size: 0.8125rem;
  }
`,k=[{title:"Get started",description:"Step-by-step guide to install and configure the necessary tools",to:"/docs/quickstart/install-guide"},{title:"Complete Guide",description:"Follow a comprehensive guide to set up your environment",to:"/docs/quickstart/complete-guide"},{title:"Check Dependencies",description:"Check the dependencies required for your setup",to:"/docs/quickstart/dependencies"},{title:"Keys Management",description:"Manage your validator keys efficiently",to:"/docs/quickstart/keys-management"},{title:"Commands",description:"Learn about advanced commands to manage or set up your node",to:"/docs/commands"}],y=[{title:"Ethereum Mainnet",description:"Deploy and manage your nodes on Ethereum Mainnet",to:"/docs/networks/mainnet"},{title:"Gnosis Network",description:"Get started with the Gnosis Mainnet and Testnet",to:"/docs/networks/gnosis"},{title:"Optimism Network",description:"Deploy on Optimism Mainnet and its test networks",to:"/docs/quickstart/running-optimism-node"},{title:"Exposing API",description:"Guide on exposing APIs for your validator setup",to:"/docs/quickstart/samples/exposing-apis"},{title:"Using Relays",description:"Learn how to utilize relays in your setup",to:"/docs/quickstart/samples/using-relays"}];function j(){return(0,x.jsx)(c,{children:(0,x.jsxs)(h,{children:[(0,x.jsxs)(g,{children:[(0,x.jsx)(w,{children:"Integrate with Sedge"}),(0,x.jsx)(u,{children:"Explore these guided tutorials to get started with Sedge for your Ethereum staking needs."}),k.map((e=>(0,x.jsxs)(f,{to:e.to,children:[(0,x.jsxs)("div",{children:[(0,x.jsx)(b,{children:e.title}),(0,x.jsx)(v,{children:e.description})]}),(0,x.jsx)(l.A,{size:20})]},e.title)))]}),(0,x.jsxs)(g,{children:[(0,x.jsx)(w,{children:"Advanced Sedge Features"}),(0,x.jsx)(u,{children:"Dive deeper into Sedge's capabilities and optimize your Ethereum staking experience."}),y.map((e=>(0,x.jsxs)(f,{to:e.to,children:[(0,x.jsxs)("div",{children:[(0,x.jsx)(b,{children:e.title}),(0,x.jsx)(v,{children:e.description})]}),(0,x.jsx)(l.A,{size:20})]},e.title)))]})]})})}var A=t(1122),z=t(6025);const S=m.A.h2`
    text-align: left;
    margin-left: 1rem;
    margin-bottom: 2rem;
    font-size: 1.5rem;

    @media (max-width: 768px) {
        font-size: 1.25rem;
        margin-bottom: 1.5rem;
    }

    @media (max-width: 480px) {
        font-size: 1.125rem;
        margin-bottom: 1rem;
        margin-left: 0.75rem;
    }
`,E=m.A.div`
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.5rem;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 1.5rem;
    width: 100%;

    @media (max-width: 1024px) {
        grid-template-columns: repeat(2, 1fr);
        gap: 1.25rem;
        padding: 0 1.25rem;
    }

    @media (max-width: 768px) {
        grid-template-columns: repeat(2, 1fr);
        gap: 1rem;
        padding: 0 1rem;
    }

    @media (max-width: 480px) {
        grid-template-columns: 1fr;
        gap: 0.75rem;
        padding: 0 0.75rem;
    }
`,L=m.A.div`
    display: flex;
    flex-direction: column;
    padding: 2rem 1.5rem 1rem;
    border: 1px solid var(--ifm-color-emphasis-300);
    border-radius: 8px;
    text-align: center;
    background-color: var(--ifm-background-surface-color);
    transition: transform 0.2s ease, box-shadow 0.2s ease;

    @media (hover: hover) {
        &:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        }
    }

    @media (max-width: 1024px) {
        padding: 1.75rem 1.25rem 0.875rem;
    }

    @media (max-width: 768px) {
        padding: 1.5rem 1rem 0.75rem;
    }

    @media (max-width: 480px) {
        padding: 1.25rem 0.875rem 0.625rem;
    }
`,q=(0,m.A)(A.A)`
    width: 80px;
    height: 80px;
    object-fit: contain;
    margin: 0 auto 0.75rem;

    @media (max-width: 1024px) {
        width: 70px;
        height: 70px;
    }

    @media (max-width: 768px) {
        width: 60px;
        height: 60px;
    }

    @media (max-width: 480px) {
        width: 50px;
        height: 50px;
        margin: 0 auto 0.5rem;
    }
`,C=m.A.h3`
    margin-bottom: 0.75rem;
    font-size: 1.1rem;
    color: var(--ifm-color-emphasis-900);

    @media (max-width: 768px) {
        font-size: 1rem;
        margin-bottom: 0.5rem;
    }

    @media (max-width: 480px) {
        font-size: 0.9375rem;
    }
`,M=m.A.a`
    display: inline-block;
    margin: 0.25rem;
    padding: 0.625rem 1rem;
    background-color: var(--ifm-color-emphasis-200);
    color: var(--ifm-color-emphasis-700);
    border-radius: 4px;
    text-decoration: none;
    font-size: 0.9rem;
    transition: background-color 0.2s ease;
    min-height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;

    &:hover {
        background-color: var(--ifm-color-emphasis-300);
        text-decoration: none;
    }

    @media (max-width: 768px) {
        padding: 0.5rem 0.875rem;
        font-size: 0.875rem;
        min-height: 36px;
    }

    @media (max-width: 480px) {
        padding: 0.4375rem 0.75rem;
        font-size: 0.8125rem;
        min-height: 32px;
    }
`,P=m.A.div`
    display: flex;
    flex-direction: column;
    align-items: stretch;
    gap: 0.5rem;

    @media (max-width: 768px) {
        gap: 0.375rem;
    }

    @media (max-width: 480px) {
        gap: 0.25rem;
    }
`,D=(0,m.A)(M)`
    display: block;
    margin: 0;
`,H=()=>(0,x.jsxs)(x.Fragment,{children:[(0,x.jsx)(S,{children:"Supported Networks"}),(0,x.jsx)(E,{children:[{name:"Base",logo:{light:"/img/chains/base-logo.png",dark:"/img/chains/base-logo.png"},networks:[{name:"Mainnet",link:"/docs/quickstart/running-optimism-node#base-support"}]},{name:"Ethereum",logo:{light:"/img/chains/eth-logo.svg",dark:"/img/chains/eth-logo.svg"},networks:[{name:"Mainnet",link:"/docs/networks/mainnet"}]},{name:"Gnosis",logo:{light:"/img/chains/gno-logo.png",dark:"/img/chains/gno-logo.png"},networks:[{name:"Mainnet",link:"/docs/networks/gnosis"}]},{name:"Optimism",logo:{light:"/img/chains/op-logo.png",dark:"/img/chains/op-logo.png"},networks:[{name:"Mainnet",link:"/docs/quickstart/running-optimism-node"}]},{name:"Testnets",networks:[{name:"Hoodi",link:"/docs/networks/hoodi"},{name:"Holesky",link:"/docs/networks/holesky"},{name:"Sepolia",link:"/docs/networks/sepolia"},{name:"Chiado",link:"/docs/networks/chiado"}]},{name:"Lido",logo:{light:"/img/chains/lido-logo.png",dark:"/img/chains/lido-logo.png"},networks:[{name:"CSM",link:"/docs/quickstart/staking-with-lido"}]}].map((e=>(0,x.jsxs)(L,{children:[e.logo&&(0,x.jsx)(q,{sources:{light:(0,z.Ay)(e.logo.light),dark:(0,z.Ay)(e.logo.dark)},alt:`${e.name} logo`}),(0,x.jsx)(C,{children:e.name}),"Testnets"===e.name?(0,x.jsx)(P,{children:e.networks.map((e=>(0,x.jsx)(D,{href:e.link,children:e.name},e.name)))}):e.networks.map((e=>(0,x.jsx)(M,{href:e.link,children:e.name},e.name)))]},e.name)))})]});var G=t(3546),N=t(5773),T=t(5404);const F=e=>{let{command:i}=e;const[t,a]=(0,r.useState)(!1),[o,n]=(0,r.useState)(!1);(0,r.useEffect)((()=>{const e=()=>{n(window.innerWidth<=480)};return e(),window.addEventListener("resize",e),()=>window.removeEventListener("resize",e)}),[]);const d=m.A.div`
        margin: 1rem 0;
        border-radius: 8px;
        overflow: hidden;
        background-color: var(--ifm-background-surface-color);
        border: 1px solid var(--ifm-color-emphasis-300);
        width: 100%;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);

        @media (max-width: 768px) {
            margin: 0.875rem 0;
        }

        @media (max-width: 480px) {
            margin: 0.75rem 0;
            border-radius: 6px;
        }
    `,s=m.A.div`
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 1rem;
        gap: 0.75rem;
        position: relative;

        @media (max-width: 768px) {
            padding: 0.875rem;
        }

        @media (max-width: 480px) {
            padding: 0.75rem 0.625rem;
        }

        @media (max-width: 768px) {
            padding: 0.875rem 1rem;
        }

        @media (max-width: 480px) {
            padding: 0.75rem;
        }
    `,p=m.A.div`
        flex: 1;
        overflow-x: auto;
        -webkit-overflow-scrolling: touch;
        position: relative;

        /* Custom scrollbar */
        ::-webkit-scrollbar {
            height: 4px;
        }

        ::-webkit-scrollbar-track {
            background: var(--ifm-color-emphasis-100);
            border-radius: 2px;
        }

        ::-webkit-scrollbar-thumb {
            background: var(--ifm-color-emphasis-300);
            border-radius: 2px;
        }

        pre {
            margin: 0;
            padding: 0;
            min-width: min-content;
        }

        code {
            font-family: var(--ifm-font-family-monospace);
            font-size: 0.9375rem;
            color: var(--ifm-color-emphasis-900);
            white-space: pre;
            padding-bottom: ${o?"0.25rem":"0"}; /* Space for scrollbar on mobile */

            @media (max-width: 768px) {
                font-size: 0.875rem;
            }

            @media (max-width: 480px) {
                font-size: 0.8125rem;
            }
        }
    `,l=m.A.button`
        display: flex;
        align-items: center;
        justify-content: center;
        background: none;
        border: none;
        padding: 0.5rem;
        cursor: pointer;
        border-radius: 4px;
        transition: all 0.2s ease;
        min-width: 36px;
        height: 36px;
        flex-shrink: 0;

        &:hover {
            background-color: var(--ifm-color-emphasis-200);
        }

        &:active {
            background-color: var(--ifm-color-emphasis-300);
            transform: scale(0.95);
        }

        svg {
            width: 18px;
            height: 18px;
            transition: transform 0.2s ease;

            @media (max-width: 480px) {
                width: 16px;
                height: 16px;
            }
        }

        @media (max-width: 480px) {
            padding: 0.375rem;
            min-width: 32px;
            height: 32px;
        }
    `;return(0,x.jsx)(d,{children:(0,x.jsxs)(s,{children:[(0,x.jsx)(p,{children:(0,x.jsx)("pre",{children:(0,x.jsx)("code",{children:i})})}),(0,x.jsx)(l,{onClick:async()=>{try{await navigator.clipboard.writeText(i),a(!0),setTimeout((()=>a(!1)),2e3)}catch(e){console.error("Failed to copy text: ",e)}},"aria-label":"Copy to clipboard",children:t?(0,x.jsx)(N.A,{style:{color:"var(--ifm-color-success)"}}):(0,x.jsx)(T.A,{style:{color:"var(--ifm-color-emphasis-600)"}})})]})})},Y=()=>{const[e,i]=r.useState("linux-macos");r.useEffect((()=>{const e=navigator.userAgent||navigator.vendor||window.opera;/windows phone/i.test(e)||/win/i.test(e)?i("windows"):(/android/i.test(e)||/iPad|iPhone|iPod/.test(e)||/mac/i.test(e)||/linux/i.test(e))&&i("linux-macos")}),[]);const t=m.A.div`
        width: 100%;
        max-width: 800px;
        margin: 0 auto;
        padding: 1.5rem;

        @media (max-width: 1024px) {
            padding: 1.25rem;
            max-width: 700px;
        }

        @media (max-width: 768px) {
            padding: 1rem;
            max-width: 600px;
        }

        @media (max-width: 480px) {
            padding: 0.75rem 0.5rem;
        }
    `,a=m.A.p`
        font-size: 1.125rem;
        margin-bottom: 1.25rem;
        color: var(--ifm-color-emphasis-800);
        text-align: center;
        max-width: 600px;
        margin-left: auto;
        margin-right: auto;
        line-height: 1.5;

        strong {
            color: var(--ifm-color-primary);
            font-weight: 600;
            white-space: nowrap;
        }

        @media (max-width: 768px) {
            font-size: 1rem;
            margin-bottom: 1rem;
            max-width: 500px;
        }

        @media (max-width: 480px) {
            font-size: 0.9375rem;
            margin-bottom: 0.875rem;
            padding: 0 0.5rem;
        }
    `;return(0,x.jsxs)(t,{children:[(0,x.jsxs)(a,{children:["To install ",(0,x.jsx)("strong",{children:"Sedge"}),", run the following command in your terminal:"]}),(0,x.jsx)(F,{command:{"linux-macos":"bash <(curl -fsSL https://github.com/NethermindEth/sedge/raw/main/scripts/install.sh)",windows:"Set-ExecutionPolicy Bypass -Scope Process -Force; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://github.com/NethermindEth/sedge/raw/main/scripts/install.ps1'))"}[e]})]})};function _(){const e=m.A.div`
        width: 100%;
        background-size: cover;
        background-position: center;
        position: relative;
        padding: 0 1rem;

        @media (max-width: 768px) {
            padding: 0 0.75rem;
        }

        @media (max-width: 480px) {
            padding: 0 0.5rem;
        }
    `,i=m.A.header`
        position: relative;
        padding: 4rem 1rem;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
        max-width: var(--ifm-container-width);
        margin: 0 auto;
        color: var(--ifm-font-color-base);

        @media (max-width: 1024px) {
            padding: 3rem 1rem;
        }

        @media (max-width: 768px) {
            padding: 2.5rem 0.75rem;
        }

        @media (max-width: 480px) {
            padding: 2rem 0.5rem;
        }
    `,t=m.A.h1`
        font-weight: 600;
        font-size: 3rem;
        margin: 1.5rem 0;
        text-align: center;
        max-width: 800px;
        line-height: 1.2;

        @media (max-width: 1024px) {
            font-size: 2.5rem;
            margin: 1.25rem 0;
        }

        @media (max-width: 768px) {
            font-size: 2rem;
            margin: 1rem 0;
        }

        @media (max-width: 480px) {
            font-size: 1.75rem;
            margin: 0.875rem 0;
        }
    `,r=m.A.div`
        margin-top: 2rem;
        width: 100%;
        max-width: 600px;
        display: flex;
        justify-content: center;

        @media (max-width: 768px) {
            margin-top: 1.5rem;
            max-width: 100%;
            padding: 0 1rem;
        }

        @media (max-width: 480px) {
            margin-top: 1.25rem;
            padding: 0 0.5rem;
        }

        /* Style the search input to be more responsive */
        :global(.navbar__search-input) {
            width: 100%;
            max-width: none;
            font-size: 1rem;
            padding: 0.75rem 1rem;

            @media (max-width: 480px) {
                font-size: 0.9375rem;
                padding: 0.625rem 0.875rem;
            }
        }
    `;return(0,x.jsx)(e,{children:(0,x.jsxs)(i,{children:[(0,x.jsx)(A.A,{style:{width:"auto",height:"auto",maxWidth:"80%",maxHeight:"120px","@media (max-width: 768px)":{maxHeight:"100px"},"@media (max-width: 480px)":{maxHeight:"80px"}},sources:{dark:(0,z.Ay)("/img/Sedge_Horizontal_Light.svg"),light:(0,z.Ay)("/img/Sedge_Horizontal_Dark.svg")},alt:"Sedge Logo"}),(0,x.jsx)(t,{children:"Easy node setup and deployment tool"}),(0,x.jsx)(Y,{}),(0,x.jsx)(r,{children:(0,x.jsx)(G.A,{})})]})})}const I=m.A.div`
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    min-height: 100vh;
    width: 100%;
    overflow-x: hidden; // Prevent horizontal scrolling
`,O=m.A.div`
    position: fixed;
    top: 0;
    left: 0;
    height: 4px;
    background-color: var(--ifm-color-primary-dark);
    z-index: 1000;
`,W=[{title:"What is Sedge",icon:d.A,to:"/docs/intro",text:"Learn about the core concepts of Sedge, its features, and how it can help you."},{title:"Get Started with Sedge",icon:s.A,to:"/docs/quickstart",text:"Learn how to install and set up Sedge for your Ethereum staking needs."},{title:"Sedge Documentation",icon:p.A,to:"/docs/quickstart/install-guide",text:"Explore the full documentation to learn about all Sedge features and capabilities."}];function $(){const e=m.A.div`
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    grid-gap: 24px;
    justify-content: center;
    margin: 0 auto;
    width: 100%;
    padding: 2rem;
    max-width: 1200px;

    @media (max-width: 1024px) {
      padding: 1.5rem;
      grid-gap: 20px;
    }

    @media (max-width: 768px) {
      padding: 1.25rem;
      grid-gap: 16px;
      grid-template-columns: 1fr;
      max-width: 600px;
    }

    @media (max-width: 480px) {
      padding: 1rem;
      grid-gap: 12px;
    }`,i=m.A.div`
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
    height: 100%;
    min-height: 240px;

    @media (max-width: 1024px) {
      padding: 1.25rem;
      min-height: 220px;
    }

    @media (max-width: 768px) {
      padding: 1.125rem;
      min-height: 200px;
    }

    @media (max-width: 480px) {
      padding: 1rem;
      min-height: 180px;
    }

    @media (hover: hover) {
      &:hover {
        transform: translateY(-5px);
        box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
      }
    }

    &:active {
      transform: translateY(1px);
    }
  `,t=m.A.h3`
    margin-bottom: 1rem;
    font-weight: 600;
    font-size: 1.25rem;

    @media (max-width: 768px) {
      font-size: 1.125rem;
      margin-bottom: 0.75rem;
    }

    @media (max-width: 480px) {
      font-size: 1rem;
      margin-bottom: 0.5rem;
    }
  `,r=m.A.p`
    margin-bottom: 1rem;
    font-weight: 400;
    font-size: 1rem;
    line-height: 1.5;

    @media (max-width: 768px) {
      font-size: 0.9375rem;
      margin-bottom: 0.75rem;
    }

    @media (max-width: 480px) {
      font-size: 0.875rem;
      margin-bottom: 0.5rem;
    }
  `,a=m.A.div`
    width: 48px;
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    background-color: var(--ifm-color-emphasis-100);
    margin-bottom: 1rem;

    @media (max-width: 768px) {
      width: 40px;
      height: 40px;
      margin-bottom: 0.75rem;
    }

    @media (max-width: 480px) {
      width: 36px;
      height: 36px;
      margin-bottom: 0.5rem;
    }
  `;return(0,x.jsx)(e,{children:W.map((e=>(0,x.jsxs)(i,{children:[(0,x.jsxs)("div",{children:[(0,x.jsx)(a,{children:(0,x.jsx)(e.icon,{size:24})}),(0,x.jsx)(t,{children:e.title}),(0,x.jsx)(r,{children:e.text})]}),(0,x.jsxs)(o.A,{to:e.to,style:{display:"flex",alignItems:"center",color:"var(--ifm-color-primary)"},children:["Learn more ",(0,x.jsx)(l.A,{size:16,style:{marginLeft:"0.5rem"}})]})]},e.title)))})}function B(){const{siteConfig:e}=(0,n.A)(),[i,t]=(0,r.useState)(0);return(0,r.useEffect)((()=>{const e=()=>{const e=document.documentElement.scrollHeight-window.innerHeight,i=window.scrollY/e*100;t(i)};return window.addEventListener("scroll",e),()=>window.removeEventListener("scroll",e)}),[]),(0,x.jsxs)(a.A,{title:`${e.title} Documentation`,description:"Technical Documentation For Sedge",children:[(0,x.jsx)(O,{style:{width:`${i}%`}}),(0,x.jsx)(_,{}),(0,x.jsxs)(I,{children:[(0,x.jsx)($,{}),(0,x.jsx)(H,{}),(0,x.jsx)(j,{})]})]})}}}]);