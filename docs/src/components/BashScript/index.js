import React, { useState, useEffect } from 'react';
import { Copy, Check } from 'lucide-react';
import styled from '@emotion/styled';

const BashCommand = ({ command }) => {
    const [copied, setCopied] = useState(false);
    const [isMobile, setIsMobile] = useState(false);

    useEffect(() => {
        const checkMobile = () => {
            setIsMobile(window.innerWidth <= 480);
        };

        checkMobile();
        window.addEventListener('resize', checkMobile);

        return () => window.removeEventListener('resize', checkMobile);
    }, []);

    const copyToClipboard = async () => {
        try {
            await navigator.clipboard.writeText(command);
            setCopied(true);
            setTimeout(() => setCopied(false), 2000);
        } catch (err) {
            console.error('Failed to copy text: ', err);
        }
    };

        const Container = styled.div`
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
    `;

    const Content = styled.div`
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
    `;

    const CodeBlock = styled.div`
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
            padding-bottom: ${isMobile ? '0.25rem' : '0'}; /* Space for scrollbar on mobile */

            @media (max-width: 768px) {
                font-size: 0.875rem;
            }

            @media (max-width: 480px) {
                font-size: 0.8125rem;
            }
        }
    `;

    const CopyButton = styled.button`
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
    `;

    return (
        <Container>
            <Content>
                <CodeBlock>
                    <pre>
                        <code>{command}</code>
                    </pre>
                </CodeBlock>
                <CopyButton
                    onClick={copyToClipboard}
                    aria-label="Copy to clipboard"
                >
                    {copied ? (
                        <Check style={{ color: 'var(--ifm-color-success)' }} />
                    ) : (
                        <Copy style={{ color: 'var(--ifm-color-emphasis-600)' }} />
                    )}
                </CopyButton>
            </Content>
        </Container>
    );
};

export default BashCommand;