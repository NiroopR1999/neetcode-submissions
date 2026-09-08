class MinStack {
    constructor() {
        this.stack=[]
    }

    /**
     * @param {number} val
     * @return {void}
     */
    push(val) {
        this.stack.push(val)
    }

    /**
     * @return {void}
     */
    pop() {
        if(this.stack.length===0) return 'underflow'
        return this.stack.pop()
    }

    /**
     * @return {number}
     */
    top() {
        if(this.stack.length===0) return 'underflow'
        return this.stack[this.stack.length-1]
    }

    /**
     * @return {number}
     */
    getMin() {
         if(this.stack.length===0) return 'underflow'
        let min=Infinity
        for(const val of this.stack)
        {
            if(val<min) min=val
        }
        return min
    }
}
