import React from 'react';
import styled from '@emotion/styled';
import ThemedImage from '@theme/ThemedImage';
import useBaseUrl from '@docusaurus/useBaseUrl';

const SectionTitle = styled.h2`
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
`;

const GridContainer = styled.div`
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
`;

const NetworkCard = styled.div`
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
`;

const NetworkLogo = styled(ThemedImage)`
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
`;

const NetworkName = styled.h3`
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
`;

const NetworkButton = styled.a`
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
`;

const TestnetsList = styled.div`
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
`;

const TestnetButton = styled(NetworkButton)`
    display: block;
    margin: 0;
`;

export const SupportedNetworks = () => {
    const networks = [
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
            name: 'Ethereum',
            logo: {
                light: '/img/chains/eth-logo.svg',
                dark: '/img/chains/eth-logo.svg'
            },
            networks: [
                { name: 'Mainnet', link: '/docs/networks/mainnet' },
            ]
        },
        {
            name: 'Gnosis',
            logo: {
                light: '/img/chains/gno-logo.png',
                dark: '/img/chains/gno-logo.png'
            },
            networks: [
                { name: 'Mainnet', link: '/docs/networks/gnosis' },
            ]
        },
        {
            name: 'Optimism',
            logo: {
                light: '/img/chains/op-logo.png',
                dark: '/img/chains/op-logo.png'
            },
            networks: [
                { name: 'Mainnet', link: '/docs/quickstart/running-optimism-node' }
            ]
        },
        {
            name: 'Testnets',
            networks: [
                { name: 'Hoodi', link: '/docs/networks/hoodi' },
                { name: 'Sepolia', link: '/docs/networks/sepolia' },
                { name: 'Chiado', link: '/docs/networks/chiado' }
            ]
        },
        {
            name: 'Lido',
            logo: {
                light: '/img/chains/lido-logo.png',
                dark: '/img/chains/lido-logo.png',
            },
            networks: [
                { name: 'CSM', link: '/docs/quickstart/staking-with-lido' }
            ]
        },
    ];

    return (
        <>
            <SectionTitle>Supported Networks</SectionTitle>
            <GridContainer>
                {networks.map((network) => (
                    <NetworkCard key={network.name}>
                        {network.logo && (
                            <NetworkLogo
                                sources={{
                                    light: useBaseUrl(network.logo.light),
                                    dark: useBaseUrl(network.logo.dark),
                                }}
                                alt={`${network.name} logo`}
                            />
                        )}
                        <NetworkName>{network.name}</NetworkName>
                        {network.name === 'Testnets' ? (
                            <TestnetsList>
                                {network.networks.map((net) => (
                                    <TestnetButton key={net.name} href={net.link}>
                                        {net.name}
                                    </TestnetButton>
                                ))}
                            </TestnetsList>
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

export default SupportedNetworks;