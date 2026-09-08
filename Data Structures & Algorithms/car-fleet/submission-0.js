class Solution {
    /**
     * @param {number} target
     * @param {number[]} position
     * @param {number[]} speed
     * @return {number}
     */
    carFleet(target, position, speed) {
        let fleet = 0;
        let fleetTime = 0;
        let cars = position.map((pos, index) => [pos, speed[index]])
            .sort((a, b) => b[0] - a[0]);


        for (const [pos, speed] of cars) {
            let timeTaken = (target - pos) / speed
            if (timeTaken > fleetTime) {
                fleet++
                fleetTime = timeTaken

            }
        }

        return fleet
    }
}