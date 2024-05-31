import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
    {
        title: 'Supporting Ethereum',
        Svg: require('@site/static/img/ethereum.svg').default,
        description: (
            <>
                We fully support Mainnet, and Testnets (Sepolia, and Holesky).

            </>
        ),
    },
    {
        title: 'Supporting Gnosis',
        Svg: require('@site/static/img/gnosis.svg').default,
        description: (
            <>
                We fully support deployment on Gnosis and Chiado(testnet).
            </>
        ),
    },
    {
        title: 'Powered by Golang',
        Svg: require('@site/static/img/go_saiyan.svg').default,
        description: (
            <>
                Go makes it easy for us to write maintainable, testable, lightweight and performant code.
            </>
        ),
    },
];

function Feature({Svg, title, description}) {
    return (
        <div className={clsx('col col--4')}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img"/>
            </div>
            <div className="text--center padding-horiz--md">
                <h3>{title}</h3>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures() {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
