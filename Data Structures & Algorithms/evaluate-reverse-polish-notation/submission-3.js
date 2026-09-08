class Solution {
    /**
     * @param {string[]} tokens
     * @return {number}
     */
    evalRPN(tokens) {
        const stack=[]

       function calculate (operator)
        {
            const operand1=stack.pop()
            const operand2=stack.pop()

            switch(operator)
            {
                case '+':
                        return operand1+operand2
                case '-':
                        return operand2-operand1
                case '*': 
                        return operand1*operand2
                case '/':
                        return Math.trunc(operand2/operand1)
                default:
                        throw new Error('Invalid Operator!!')
            }
        }

        for(const token of tokens)
        {
            if(!isNaN(token))
            {
                stack.push(Number(token))
            }
            else{
                stack.push(calculate(token))
            }
        }
        return stack.pop()
    }
}
