import React, { useState, useEffect } from 'react';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styled from '@emotion/styled';
import { ArrowUpRight as LinkIcon, BookOpen, HelpCircle, Info } from 'react-feather';
import {QuickLinks} from "../components/QuickLinks";
import {SupportedNetworks} from "../components/SupportedNetworks";
import {Header} from "../components/Header";


const Container = styled.div`
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    min-height: 100vh;
    width: 100%;
    overflow-x: hidden; // Prevent horizontal scrolling
`;


const ProgressBar = styled.div`
    position: fixed;
    top: 0;
    left: 0;
    height: 4px;
    background-color: var(--ifm-color-primary-dark);
    z-index: 1000;
`;




const actions = [
    {
        title: 'What is Sedge',
        icon: Info,
        to: '/docs/intro',
        text: 'Learn about the core concepts of Sedge, its features, and how it can help you.',
    },
    {
        title: 'Get Started with Sedge',
        icon: HelpCircle,
        to: '/docs/quickstart',
        text: 'Learn how to install and set up Sedge for your Ethereum staking needs.',
    },
    {
        title: 'Sedge Documentation',
        icon: BookOpen,
        to: '/docs/quickstart/install-guide',
        text: 'Explore the full documentation to learn about all Sedge features and capabilities.',
    },
];



function Content() {
    const Row = styled.div`
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
    }`;

    const Card = styled.div`
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
  `;

    const CardTitle = styled.h3`
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
  `;

    const CardText = styled.p`
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
  `;

    const IconWrapper = styled.div`
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
  `;

    return (
        <Row>
            {actions.map((action) => (
                <Card key={action.title}>
                    <div>
                        <IconWrapper>
                            <action.icon size={24} />
                        </IconWrapper>
                        <CardTitle>{action.title}</CardTitle>
                        <CardText>{action.text}</CardText>
                    </div>
                    <Link to={action.to} style={{ display: 'flex', alignItems: 'center', color: 'var(--ifm-color-primary)' }}>
                        Learn more <LinkIcon size={16} style={{ marginLeft: '0.5rem' }} />
                    </Link>
                </Card>
            ))}
        </Row>
    );
}



export default function Home() {
    const { siteConfig } = useDocusaurusContext();
    const [scrollProgress, setScrollProgress] = useState(0);

    useEffect(() => {
        const handleScroll = () => {
            const totalHeight = document.documentElement.scrollHeight - window.innerHeight;
            const progress = (window.scrollY / totalHeight) * 100;
            setScrollProgress(progress);
        };

        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, []);

    return (
        <Layout
            title={`${siteConfig.title} Documentation`}
            description="Technical Documentation For Sedge">
            <ProgressBar style={{ width: `${scrollProgress}%` }} />
            <Header />
            <Container>
                <Content />
                <SupportedNetworks />
                <QuickLinks />
            </Container>
        </Layout>
    );
}
