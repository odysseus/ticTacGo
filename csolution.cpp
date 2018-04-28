#include <iostream>

using namespace std;

char mark[] = {'X', 'O'};
int x_win = 0, o_win = 0, draw = 0;

bool win(char board[3][3])
{
	for (int i = 0; i < 3; i++) {
		if (board[i][0] != '\0' && board[i][0] == board[i][1] && board[i][2] == board[i][1]) {
			return true;
		}
		if (board[0][i] != '\0' && board[0][i] == board[1][i] && board[2][i] == board[1][i]) {
			return true;
		}
	}
	
	if (board[0][0] != '\0' && board[0][0] == board[1][1] && board[1][1] == board[2][2]) {
		return true;
	}
	
	if (board[0][2] != '\0' && board[0][2] == board[1][1] && board[1][1] == board[2][0]) {
		return true;
	}
	
	return false;
}

int move(char board[3][3], int turn, int chances)
{
	if (win(board)) {
		if(turn == 0) {
			o_win++;
		} else {
			x_win++;
		}
		return 1;
	}
	
	if (chances == 9) {
		draw++;
		return 1;
	}
	
	int ans = 0;
	for (int i = 0; i < 3; i++) {
		for (int j = 0; j < 3; j++) {
			if (board[i][j] == '\0') {
				board[i][j] = mark[turn];
				ans += move(board, 1 - turn, chances + 1);
				board[i][j] = '\0';
			}
		}
	}
	
	return ans;
}

int main() {
	char board[3][3] = {'\0', '\0', '\0', '\0', '\0', '\0', '\0', '\0', '\0'};
	
	int ans = move(board, 0, 0);
	
	cout << "Total Games: " << ans << endl;
	cout << "X Wins: " << x_win << endl;
	cout << "Y wins: " << o_win << endl;
	cout << "Draw: " << draw << endl;
	
	return 0;
}