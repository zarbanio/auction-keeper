import {GoerliConfig} from "./uniswap-v3-callee/goerli";
import {MainConfig} from "./uniswap-v3-callee/main";


export const GetConfig = (contract: string, network: string): any => {
    if (contract == "uniswapV3Callee") {
        if (network == "goerli") {
            return GoerliConfig
        }
        return MainConfig
    }
}