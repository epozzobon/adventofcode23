#include <fstream>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

int main()
{
    string line;
    ifstream myfile("input.txt");
    if (!myfile.is_open()) {
        cout << "Unable to open file" << endl;
        return 1;
    }

    vector<string> textDigits = { "one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine" };

    unsigned sum = 0;
    while (getline(myfile, line)) {
        unsigned firstDigit = 0;
        unsigned lastDigit = 0;
        for (unsigned i = 0; i < line.length(); i++) {
            unsigned digit = 0;
            if (string::npos != string("123456789").find(line[i])) {
                digit = line[i] - '0';
            }

            for (unsigned j = 0; j < textDigits.size(); j++) {
                if (0 == line.substr(i, textDigits[j].length()).compare(textDigits[j])) {
                    digit = j + 1;
                }
            }

            if (digit != 0) {
                if (0 == firstDigit) {
                    firstDigit = digit;
                }
                lastDigit = digit;
            }
        }
        unsigned number = firstDigit * 10 + lastDigit;
        sum += number;
        cout << firstDigit << lastDigit << " " << number << endl;
    }
    cout << "Sum is " << sum << endl;
    myfile.close();
    return 0;
}