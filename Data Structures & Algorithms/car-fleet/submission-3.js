class Solution {
    /**
     * @param {number} target
     * @param {number[]} position
     * @param {number[]} speed
     * @return {number}
     */
    carFleet(target, position, speed) {
        const cars=position.map((pos,index)=>[pos,speed[index]]).sort((a,b)=>b[0]-a[0])
        let fleet=0
        let fleetTime=0

        for(const [pos,speed] of cars)
        {
            const timeTaken=(target-pos)/speed

            if(timeTaken>fleetTime)
            {
                fleet++
                fleetTime=timeTaken
            }
        }
        return fleet
    }
}
