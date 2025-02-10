import React from 'react';
import BashCommand from '../BashScript';
import styled from '@emotion/styled';

const InstallCommand = () => {
    const [os, setOs] = React.useState('linux-macos');

    React.useEffect(() => {
        const userAgent = navigator.userAgent || navigator.vendor || window.opera;

        if (/windows phone/i.test(userAgent) || /win/i.test(userAgent)) {
            setOs('windows');
        } else if (/android/i.test(userAgent)) {
            setOs('linux-macos');
        } else if (/iPad|iPhone|iPod/.test(userAgent) || /mac/i.test(userAgent)) {
            setOs('linux-macos');
        } else if (/linux/i.test(userAgent)) {
            setOs('linux-macos');
        }
    }, []);

    const commands = {
        'linux-macos': "bash <(curl -fsSL https://github.com/NethermindEth/sedge/raw/main/scripts/install.sh)",
        windows: "Set-ExecutionPolicy Bypass -Scope Process -Force; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://github.com/NethermindEth/sedge/raw/main/scripts/install.ps1'))",
    };

    const Container = styled.div`
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
    `;

    const Description = styled.p`
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
    `;

    return (
        <Container>
            <Description>
                To install <strong>Sedge</strong>, run the following command in your terminal:
            </Description>
            <BashCommand command={commands[os]} />
        </Container>
    );
};

export default InstallCommand;