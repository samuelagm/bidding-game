//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract BiddingWar is Ownable {
    uint256 private _accumulatedCommission;
    address private _currentWinner;
    uint256 private _winningBid;
    uint8 private immutable _commison = 5;
    uint256 private _round;
    uint8 private _step;
    uint256 private _roundEndTime;

    error activeBidRound(string message);
    event NewBidRound(uint256 round, uint256 expiry);
    event NewBid(
        uint256 round,
        address bidder,
        uint256 bidAmount,
        uint256 timeLeft
    );
    event CashOut(uint256 round, address winner, uint256 amount);

    struct BidRound {
        address winner;
        uint256 winningBid;
        uint256 accumulatedBids;
    }

    mapping(uint256 => BidRound) private _winners;

    modifier afterBidWar() {
        if (block.timestamp <= _roundEndTime) {
            revert activeBidRound("a bidding war is ongoing");
        }
        _;
    }

    constructor() {
        _step = 0;
        _round = 0;
        _roundEndTime = block.timestamp + 60 minutes;
        _accumulatedCommission = 0;
        emit NewBidRound(_round, _roundEndTime);
    }

    function bid() public payable {
        require(block.timestamp <= _roundEndTime, "bid round has ended");
        require(msg.value > 0, "invalid bid");

        if (_step == 0) {
            // Round Initialization
            _step += 1;
            setVars(msg.sender, msg.value);
            _winners[_round] = BidRound(
                _currentWinner,
                _winningBid,
                _winningBid
            );
            announce();
            return;
        }

        require(
            msg.value > _winners[_round].winningBid,
            "bid lower than current bid"
        );

        setVars(msg.sender, msg.value);
        _winners[_round] = BidRound(
            _currentWinner,
            _winningBid,
            _winningBid + _winners[_round].accumulatedBids
        );
        announce();
    }

    function setVars(address winner, uint256 bidAmount) internal {
        uint256 commisionPercent = bidAmount * (_commison / 100);
        _winningBid = bidAmount - commisionPercent;
        _accumulatedCommission += commisionPercent;
        _currentWinner = winner;
    }

    function announce() internal {
        _roundEndTime += 10 minutes;
        emit NewBid(
            _round,
            _currentWinner,
            _winningBid,
            _roundEndTime - block.timestamp
        );
    }

    function restart() public onlyOwner afterBidWar {
        _roundEndTime = block.timestamp + 60 minutes;
        _round += 1;
        emit NewBidRound(_round, _roundEndTime);
    }

    function payWinner(uint256 round) public onlyOwner afterBidWar {
        BidRound memory roundInfo = _winners[round];
        require(roundInfo.accumulatedBids > 0, "winner already paid");
        _winners[round].accumulatedBids = 0;
        (bool success, ) = roundInfo.winner.call{
            value: roundInfo.accumulatedBids
        }("");
        require(success, "unable to pay winner");
        emit CashOut(round, roundInfo.winner, roundInfo.accumulatedBids);
    }

    function getCurrentRoundNumber() public view returns (uint256) {
        return _round;
    }

    function getLastBid() public view returns (BidRound memory) {
        return _winners[_round];
    }

    function getBidAtRound(uint256 round)
        public
        view
        returns (BidRound memory)
    {
        require(round <= _round, "invalid round");
        return _winners[round];
    }

    function getCommissions() public view returns (uint256) {
        return _accumulatedCommission;
    }

    function gameIsRunning() public view returns (bool) {
        if (block.timestamp <= _roundEndTime){
            return true;
        return false;
    }

    function withdraw() public onlyOwner {
        (bool success, ) = msg.sender.call{value: _accumulatedCommission}("");
        require(success, "unable to withdraw commissions");
    }
}
