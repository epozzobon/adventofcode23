#include <assert.h>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;

vector<string> split(string s, string delimiter)
{
    size_t pos_start = 0, pos_end, delim_len = delimiter.length();
    string token;
    vector<string> res;

    while ((pos_end = s.find(delimiter, pos_start)) != string::npos) {
        token = s.substr(pos_start, pos_end - pos_start);
        pos_start = pos_end + delim_len;
        res.push_back(token);
    }

    res.push_back(s.substr(pos_start));
    return res;
}

vector<int> splitInts(string s)
{
    vector<int> output = {};
    vector<string> pieces = split(s, " ");
    for (unsigned i = 0; i < pieces.size(); i++) {
        if (pieces[i].length() > 0) {
            output.push_back(stoi(pieces[i]));
        }
    }
    return output;
}

unsigned countWaysToBeat(unsigned long long raceDuration, unsigned long long distanceRecord) {
	unsigned long long minimumSpeed = 0;
	for (unsigned long long speed = 0; speed < raceDuration; speed++) {
		if (raceDuration*speed-speed*speed > distanceRecord) {
			minimumSpeed = speed;
			break;
		}
	}

	unsigned long long maximumSpeed = 0;
	for (unsigned long long speed = raceDuration - 1; speed >= 0; speed--) {
		if (raceDuration*speed-speed*speed > distanceRecord) {
			maximumSpeed = speed;
			break;
		}
	}
	unsigned long long waysToBeat = maximumSpeed - minimumSpeed + 1;
	cout << minimumSpeed << " " <<  maximumSpeed << endl;
	cout << waysToBeat << endl;
	return waysToBeat;
}

int main()
{
    string line;
    ifstream myfile("input.txt");
    if (!myfile.is_open()) {
        cout << "Unable to open file" << endl;
        return 1;
    }

    getline(myfile, line);
    vector<string> timePieces = split(line, ":");
    assert(timePieces.size() == 2);
    vector<int> times = splitInts(timePieces[1]);

    getline(myfile, line);
    vector<string> distancePieces = split(line, ":");
    assert(distancePieces.size() == 2);
    vector<int> distances = splitInts(distancePieces[1]);

    unsigned output = 1;
    for (unsigned i = 0; i < times.size(); i++) {
        int raceDuration = times[i];
        int distanceRecord = distances[i];

        cout << raceDuration << " " << distanceRecord << endl;
        int waysToBeat = countWaysToBeat(raceDuration, distanceRecord);
        output *= waysToBeat;
    }
    cout << "Part1 solution: " << output << endl;

    string timeString = "";
    for (string s : split(timePieces[1], " ")) {
        timeString += s;
    }
    string distanceString = "";
    for (string s : split(distancePieces[1], " ")) {
        distanceString += s;
    }

    unsigned long long time = stoull(timeString);
    unsigned long long distance = stoull(distanceString);
    cout << time << " " << distance << endl;
    unsigned waysToBeat = countWaysToBeat(time, distance);
    cout << "Part2 solution: " << waysToBeat << endl;

    return 0;
}