import React from 'react';
import styled from '@emotion/styled';
import ThemedImage from '@theme/ThemedImage';
import useBaseUrl from '@docusaurus/useBaseUrl';

const SectionTitle = styled.h2`
    text-align: left;
    margin-left: 1rem;
    margin-bottom: 2rem;
    font-size: 1.5rem;
`;

const GridContainer = styled.div`
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 1rem;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 1rem;
`;

const NetworkCard = styled.div`
    padding: 1rem;
    border: 1px solid var(--ifm-color-emphasis-300);
    border-radius: 8px;
    text-align: center;
    background-color: var(--ifm-background-surface-color);
`;

const NetworkLogo = styled(ThemedImage)`
    width: 80px;
    height: 80px;
    object-fit: contain;
    margin-bottom: 0.5rem;
`;

const NetworkName = styled.h3`
    margin-bottom: 0.5rem;
    font-size: 1.1rem;
    color: var(--ifm-color-emphasis-900);
`;

const NetworkButton = styled.a`
    display: inline-block;
    margin: 0.25rem;
    padding: 0.5rem 1rem;
    background-color: var(--ifm-color-emphasis-200);
    color: var(--ifm-color-emphasis-700);
    border-radius: 4px;
    text-decoration: none;
    font-size: 0.9rem;
`;

const ComingSoonLabel = styled.span`
    display: inline-block;
    margin-top: 0.5rem;
    padding: 0.25rem 0.5rem;
    background-color: var(--ifm-color-emphasis-100);
    color: var(--ifm-color-emphasis-600);
    border-radius: 4px;
    font-size: 0.9rem;
`;

export const SupportedNetworks = () => {
    const networks = [
        {
            name: 'Ethereum',
            logo: {
                light: '/img/chains/eth-logo.svg',
                dark: '/img/chains/eth-logo.svg'
            },
            networks: [
                { name: 'Mainnet', link: '/networks/mainnet' },
                { name: 'Testnet', link: '/networks/sepolia' },
            ]
        },
        {
            name: 'Gnosis',
            logo: {
                light: '/img/chains/gno-logo.png',
                dark: '/img/chains/gno-logo.png'
            },
            networks: [
                { name: 'Mainnet', link: '/networks/gnosis' },
                { name: 'Chiado', link: '/networks/chiado' }
            ]
        },
        {
            name: 'Optimism',
            logo: {
                light: '/img/chains/op-logo.png',
                dark: '/img/chains/op-logo.png'
            },
            networks: [
                { name: 'Mainnet', link: 'docs/quickstart/running-optimism-node' }
            ]
        },
        {
            name: 'Base',
            logo: {
                light: '/img/chains/base-logo.png',
                dark: '/img/chains/base-logo.png',
            },
            networks: [
                { name: 'Mainnet', link: '/docs/quickstart/running-optimism-node#base-support' },
            ]
        },
        {
            name: 'Obol',
            logo: {
                light: '/img/chains/obol-logo.png',
                dark: '/img/chains/obol-logo.png',
            },
            comingSoon: true
        },
    ];

    return (
        <>
            <SectionTitle>Supported Networks</SectionTitle>
            <GridContainer>
                {networks.map((network) => (
                    <NetworkCard key={network.name}>
                        <NetworkLogo
                            sources={{
                                light: useBaseUrl(network.logo.light),
                                dark: useBaseUrl(network.logo.dark),
                            }}
                            alt={`${network.name} logo`}
                        />
                        <NetworkName>{network.name}</NetworkName>
                        {network.comingSoon ? (
                            <ComingSoonLabel>Coming Soon</ComingSoonLabel>
                        ) : (
                            network.networks.map((net) => (
                                <NetworkButton key={net.name} href={net.link}>
                                    {net.name}
                                </NetworkButton>
                            ))
                        )}
                    </NetworkCard>
                ))}
            </GridContainer>
        </>
    );
};