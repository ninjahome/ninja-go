//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.11;

import "./owned.sol";
import "./safemath.sol";
import "./ninja-token.sol";


contract NinjaChatLicense is owned{
    using SafeMath for uint256;
    NinjaToken public token;
    address public ninjaAddr;
    address[] public WhiteLists;

    struct LicenseData {
         bool used;
         uint32 nDays;
    }
    //issue addr, random id,
    mapping(address=>mapping(bytes32=>LicenseData)) public Licenses;

    struct UserData {
        uint64 EndDays;
        uint32 TotalCoins;
    }
    //user id,
    mapping(bytes32=>UserData) public UserLicenses;

    event GenerateLicenseEvent(
        address indexed issueAddr,
        bytes32 id,
        uint32  nDays
    );

    event BindLicenseEvent(address indexed issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays);

    event ChargeUserEvent(address indexed payerAddr, bytes32 userAddr, uint32 nDays);

    event TransferLicenseEvent(address indexed executeAddr, bytes32 from, bytes32 to, uint32 nDays);

    constructor(address tAddr, address nAddr) {
        token = NinjaToken(tAddr);
        ninjaAddr = nAddr;
    }

    function AddWhiteListAddress(address executeAddr) external onlyOwner{
        for(uint i=0;i<WhiteLists.length;i++){
            if (WhiteLists[i] == executeAddr){
                revert("address already in whitelist");
            }
        }

        WhiteLists.push(executeAddr);
    }

    function DelWhiteListAddress(address executeAddr) external onlyOwner{
        uint idx = 0x00FFFFFF;

        for(uint i=0;i<WhiteLists.length;i++){
            if (WhiteLists[i] == executeAddr){
                idx = i;
                break;
            }
        }

        if(idx >= 0x00FFFFFF){
            revert("address not found");
        }
        WhiteLists[idx]=WhiteLists[WhiteLists.length-1];
        WhiteLists.pop();
    }


    function SetTokenAddr(address tAddr) external onlyOwner{
        token = NinjaToken(tAddr);
    }

    function SetNinjaAddr(address nAddr) external onlyOwner{
        ninjaAddr = nAddr;
    }

    function Setting(address tAddr, address nAddr) external onlyOwner{
        token = NinjaToken(tAddr);
        ninjaAddr = nAddr;
    }

    function GetSettings() external view returns(address, address){
        return (address(token),ninjaAddr);
    }

    function GenerateLicense(bytes32 id, uint32 nDays) external {

        require(nDays > 0,"time must large than 0");

        LicenseData memory ld = Licenses[msg.sender][id];
        require(ld.nDays == 0, "id is used");

        token.transferFrom(msg.sender,ninjaAddr, nDays*10**(token.decimals()));

        Licenses[msg.sender][id] = LicenseData(false, nDays);

        emit GenerateLicenseEvent(msg.sender, id, nDays);
    }

    function GetUserLicense(bytes32 userAddr) external view returns (uint64, uint32){
        UserData memory ud = UserLicenses[userAddr];

        return (ud.EndDays,ud.TotalCoins);
    }

    function ChargeUser(bytes32 userAddr, uint32 nDays) external{
         require(nDays > 0,"time must large than 0");

         token.transferFrom(msg.sender,ninjaAddr, nDays*10**(token.decimals()));

         UserData memory ud = UserLicenses[userAddr];

         uint curTime = block.timestamp;

         if (curTime  > ud.EndDays){
             UserLicenses[userAddr] = UserData(uint64(curTime+(3600*24*nDays)),ud.TotalCoins+nDays);
         }else{
             UserLicenses[userAddr] = UserData(uint64(ud.EndDays+(3600*24*nDays)),ud.TotalCoins+nDays);
         }

         emit ChargeUserEvent(msg.sender, userAddr, nDays);
    }

    function BindLicense(address issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays, bytes memory signature) external{
        LicenseData memory ld = Licenses[issueAddr][id];
        require(ld.used == false, "id is used");
        require(ld.nDays == nDays, "nDays not matched");

        bytes32 message = keccak256(abi.encode(this,issueAddr, id, nDays));
        bytes32 msgHash = prefixed(message);
        require(recoverSigner(msgHash, signature) == issueAddr);

        Licenses[issueAddr][id] = LicenseData(true, ld.nDays);

        UserData memory ud = UserLicenses[recvAddr];

        uint curTime = block.timestamp;

        if (curTime  > ud.EndDays){
            UserLicenses[recvAddr] = UserData(uint64(curTime+(86400*nDays)),ud.TotalCoins+nDays);
        }else{
            UserLicenses[recvAddr] = UserData(uint64(ud.EndDays+(86400*nDays)),ud.TotalCoins+nDays);
        }

        emit BindLicenseEvent(issueAddr, recvAddr, id, nDays);
    }

    function TransferLicense(bytes32 from, bytes32 to, uint32 nDays) external{
        bool find = false;
        for (uint i=0;i<WhiteLists.length;i++){
            if (WhiteLists[i] == msg.sender){
                find = true;
                break;
            }
        }
        require(find == true,"not a valid address");

        UserData memory ud = UserLicenses[from];

        uint curTime = block.timestamp;
        uint udnDays = (ud.EndDays - curTime)/86400;

        require(udnDays > nDays,"day time not enough");

        UserLicenses[from] = UserData(ud.EndDays-(nDays*86400),ud.TotalCoins-nDays);

        if (curTime  >  UserLicenses[to].EndDays){
            UserLicenses[to] = UserData(uint64(curTime+(86400*nDays)),ud.TotalCoins+nDays);
        }else{
            UserLicenses[to] = UserData(uint64(ud.EndDays+(86400*nDays)),ud.TotalCoins+nDays);
        }
        emit TransferLicenseEvent(msg.sender, from, to, nDays);
    }


    function prefixed(bytes32 hash) internal pure returns (bytes32) {
       return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
    function recoverSigner(bytes32 message, bytes memory sig) internal pure  returns (address) {
       (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);
       return ecrecover(message, v, r, s);
    }
    /// signature methods.
    function splitSignature(bytes memory sig) internal pure returns (uint8 v, bytes32 r, bytes32 s) {
       require(sig.length == 65);
       assembly {
       // first 32 bytes, after the length prefix.
           r := mload(add(sig, 32))
       // second 32 bytes.
           s := mload(add(sig, 64))
       // final byte (first byte of the next 32 bytes).
           v := byte(0, mload(add(sig, 96)))
       }
       return (v, r, s);
    }

}
