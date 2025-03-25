import React from "react";
import styled from "@emotion/styled";
import ThemedImage from "@theme/ThemedImage";
import useBaseUrl from "@docusaurus/useBaseUrl";
import SearchBar from "@theme/SearchBar";
import InstallCommand from "../InstallCommand";

export function Header() {
    const HeaderWrapper = styled.div`
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
    `;


    const HeaderContent = styled.header`
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
    `;

    const Title = styled.h1`
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
    `;

    const SearchWrapper = styled.div`
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
    `;

    return (
        <HeaderWrapper>
            <HeaderContent>
                <ThemedImage
                    style={{
                        width: 'auto',
                        height: 'auto',
                        maxWidth: '80%',
                        maxHeight: '120px',
                        '@media (max-width: 768px)': {
                            maxHeight: '100px',
                        },
                        '@media (max-width: 480px)': {
                            maxHeight: '80px',
                        },
                    }}
                    sources={{
                        dark: useBaseUrl('/img/Sedge_Horizontal_Light.svg'),
                        light: useBaseUrl('/img/Sedge_Horizontal_Dark.svg'),
                    }}
                    alt="Sedge Logo"
                />
                <Title>Easy node setup and deployment tool</Title>
                <InstallCommand/>
                <SearchWrapper>
                    <SearchBar />
                </SearchWrapper>
            </HeaderContent>
        </HeaderWrapper>
    );
}