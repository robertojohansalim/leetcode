#include <iostream>
using namespace std;

class Solution {
public:
    int buyChoco(vector<int>& prices, int money) {
        return naiveSolution(prices, money);
    }


    int naiveSolution(vector<int>& prices, int money){
        int maximum = 0;
        bool assigned = false;
        for(int i = 0; i < prices.size(); i++){
            for(int j = i + 1; j < prices.size(); j++){
                int leftover = money - (prices[i] + prices[j]);

                if (leftover < 0){
                    continue;
                }

                if (leftover >= maximum){
                    maximum = leftover;
                    assigned = true;
                }
            }
        }
        if (!assigned) return money;

        return maximum;
    }
};


int main(int argv, char*[]) 
{
    vector<int> vect = {1, 2, 2}; 
    // vector<int> vect = {10, 20, 30};

    Solution s;
    int result =  s.buyChoco(vect, 3);
    cout << result << "\n";
}     



