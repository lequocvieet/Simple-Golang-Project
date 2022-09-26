// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

contract Album {
    struct AlbumInfo {
        int256 price;
        string creater;
        string title;
    }
    AlbumInfo[] public albumInfo;

    function setInfo(
        int256 _price,
        string memory _title,
        string memory _creater
    ) public {
        albumInfo.push(AlbumInfo(_price,_title,_creater));
    }

    function getInfo(uint256 index) public view returns (AlbumInfo memory) {
        return albumInfo[index];
    }
}
