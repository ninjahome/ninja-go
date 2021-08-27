//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.11;

import "./owned.sol";
import "./safemath.sol";



contract NinjaConfig is owned{
    using SafeMath for uint256;

    struct LicenseConfig {
        address tokenAddr;
        address licenseContractAddr;
        bytes  accessAddr;
    }

    struct BootsTrap{
        bytes4 ipAddr;
        uint16[6] port;
    }

    LicenseConfig public lCfg;
    BootsTrap[]  public boots;

    constructor(){

    }

    event AddBootsTrapEvent(bytes4 ipAddr, uint16 port1, uint16 port2,uint16 port3, uint16 port4,uint16 port5, uint16 port6);
    event DelBootsTrapEvent(bytes4 ipAddr);

    function LicenseConfigSet(address tAddr, address lcAddr,bytes memory aAddr) external onlyOwner{
        require(tAddr != address(0),"token address not correct");
        require(lcAddr != address(0),"license contract address not correct");

        lCfg.tokenAddr = tAddr;
        lCfg.licenseContractAddr = lcAddr;
        lCfg.accessAddr = aAddr;
    }

    function GetLicenseConfig() external view returns (address,address,bytes memory){
        return (lCfg.tokenAddr,lCfg.licenseContractAddr,lCfg.accessAddr);
    }

    function AddBootsTrap(bytes4 ipAddr,uint16 port1, uint16 port2,uint16 port3, uint16 port4,uint16 port5, uint16 port6) external onlyOwner{
        for (uint i=0;i<boots.length;i++){
            if (boots[i].ipAddr == ipAddr){
                revert("ip Address already in bootstrap list");
            }
        }

        boots.push(BootsTrap(ipAddr,[port1,port2,port3,port4,port5,port6]));

        emit AddBootsTrapEvent(ipAddr,port1, port2,port3, port4,port5, port6);
    }

    function DelBootsTrap(bytes4 ipAddr) external onlyOwner{
        uint idx = 0x00FFFFFF;
        for(uint i=0;i<boots.length; i++){
            if( boots[i].ipAddr == ipAddr){
                idx = i;
                break;
            }
        }
        if (idx >= 0x00FFFFFF){
            revert("ip address not found");
        }
        boots[idx] = boots[boots.length-1];
        boots.pop();

        emit DelBootsTrapEvent(ipAddr);
    }

    function GetIpAddrList() external view returns (bytes4[] memory){
        bytes4[] memory addrList = new bytes4[](32);
        for(uint i=0;i<boots.length;i++){
            addrList[i]=boots[i].ipAddr;
        }

        return (addrList);
    }

    function GetIPPort(bytes4 ipAddr) external view returns(uint16,uint16,uint16,uint16,uint16,uint16){
        for(uint i=0;i<boots.length;i++){
            if (boots[i].ipAddr == ipAddr){
                return (boots[i].port[0],boots[i].port[1],boots[i].port[2],boots[i].port[3],boots[i].port[4],boots[i].port[5]);
            }
        }
        revert("no find boots trap");
    }

    function GetIPPortByIdx(uint idx) external view returns(uint16,uint16,uint16,uint16,uint16,uint16){

        if(idx >= boots.length){
            revert("not find boots trap index");
        }
        return (boots[idx].port[0],boots[idx].port[1],boots[idx].port[2],boots[idx].port[3],boots[idx].port[4],boots[idx].port[5]);
    }

}

