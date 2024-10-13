import React, { useState } from 'react';
import { Copy, Check } from 'lucide-react';

const BashCommand = () => {
    const [copied, setCopied] = useState(false);
    const command = "curl -sL brew install nethermindeth/sedge/sedge";

    const copyToClipboard = async () => {
        try {
            await navigator.clipboard.writeText(command);
            setCopied(true);
            setTimeout(() => setCopied(false), 2000);
        } catch (err) {
            console.error('Failed to copy text: ', err);
        }
    };

    return (
        <div className="rounded-md bg-gray-900 text-white font-mono text-sm">
            <div className="flex items-center p-4">
                <div className="flex-grow overflow-x-auto mr-4">
                    <pre className="flex items-center">
                        <code className="overflow-x-auto">{command}</code>
                    </pre>
                </div>
                <div className="flex-shrink-0">
                    <button
                        onClick={copyToClipboard}
                        className="p-2 rounded hover:bg-gray-700 transition-colors"
                        aria-label="Copy to clipboard"
                    >
                        {copied ? (
                            <Check className="h-5 w-5 text-green-400"/>
                        ) : (
                            <Copy className="h-5 w-5 text-gray-400"/>
                        )}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default BashCommand;