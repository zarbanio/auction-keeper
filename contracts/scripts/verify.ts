const hre = require("hardhat");


const fatalErrors = [
    `The address provided as argument contains a contract, but its bytecode`,
    `Daily limit of 100 source code submissions reached`,
    `has no bytecode. Is the contract deployed to this network`,
    `The constructor for`,
];

const okErrors = [`Contract source code already verified`];

const unableVerifyError = 'Fail - Unable to verify';

function delay(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

export const verifyEtherscanContract = async (contractAddress: string, params: any) => {
    try {
        console.log(
            '[ETHERSCAN][WARNING] Verifying deployed contract'
        );
        const msDelay = 3000;
        const times = 4;

        await verifyWithRetry(contractAddress, params, times, msDelay);
    } catch (error) {}
}

const verifyWithRetry = async (contractAddress: string, params: any, times: number, msDelay: number) => {
    let counter = times;
    await delay(msDelay);

    try {
        if (times > 1) {
            await verify(contractAddress, params)
        } else if (times === 1) {
            console.log('[ETHERSCAN][WARNING] Trying to verify via uploading all sources.');
            await verify(contractAddress, params)
        } else {
            console.error(
                '[ETHERSCAN][ERROR] Errors after all the retries, check the logs for more information.'
            );
        }
    } catch (error) {
        counter--;

        if (okErrors.some((okReason) => (error as Error).message.includes(okReason))) {
            console.info('[ETHERSCAN][INFO] Skipping due OK response: ', (error as Error).message);
            return;
        }

        if (fatalErrors.some((fatalError) => (error as Error).message.includes(fatalError))) {
            console.error(
                '[ETHERSCAN][ERROR] Fatal error detected, skip retries and resume deployment.',
                (error as Error).message
            );
            return;
        }
        console.error('[ETHERSCAN][ERROR]', (error as Error).message);
        console.log();
        console.info(`[ETHERSCAN][[INFO] Retrying attemps: ${counter}.`);
        if ((error as Error).message.includes(unableVerifyError)) {
            console.log('[ETHERSCAN][WARNING] Trying to verify via uploading all sources.');
            delete params.relatedSources;
        }
        await verifyWithRetry(contractAddress, params, counter, msDelay);
    }
}

const verify = async (address: string, params: any) => {
    return hre.run("verify:verify", {
        address: address,
        constructorArguments: params,
    })
}