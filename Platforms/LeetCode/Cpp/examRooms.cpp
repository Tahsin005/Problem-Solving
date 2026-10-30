class ExamRoom {
private:
    set<int> seats;
	int capacity;
public:
    ExamRoom(int n) {
        capacity = n;
    }
    
    int seat() {
        int dist = 0;
        int curr = 0;

        if (!seats.empty()) {
            auto itr = seats.begin();

            dist = *itr;
            if (dist == 0) {
                itr++;
            }

            while (itr != seats.end()) {
                int mid_dist = (*itr - *prev(itr)) / 2;
				if (dist < mid_dist) {
					dist = mid_dist;
					curr = *prev(itr) + dist;
				}
				itr++;
			}
            
			if (dist < ((capacity - 1) - *(seats.rbegin()))) {
				curr = capacity - 1;
			}
        }

        return *(seats.insert(curr).first);
    }
    
    void leave(int p) {
        seats.erase(p);
    }
};

/**
 * Your ExamRoom object will be instantiated and called as such:
 * ExamRoom* obj = new ExamRoom(n);
 * int param_1 = obj->seat();
 * obj->leave(p);
 */
