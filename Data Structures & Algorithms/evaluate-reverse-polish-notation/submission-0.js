class Solution {
    /**
     * @param {string[]} tokens
     * @return {number}
     */
    evalRPN(tokens) {
        let res = [];
        
        function calculate(operator) {
            const operand2 = res.pop();
            const operand1 = res.pop();
            
            switch (operator) {
                case '+':
                    return operand1 + operand2;
                case '-':
                    return operand1 - operand2;
                case '*':
                    return operand1 * operand2;
                case '/':
                    return Math.trunc(operand1 / operand2);
                default:
                    throw new Error('Invalid operator');
            }
        }
        
        for (const token of tokens) {
            if (!isNaN(token)) {
                res.push(Number(token));
            } else {
                res.push(calculate(token));
            }
        }
        
        return res.pop();
    }
}
