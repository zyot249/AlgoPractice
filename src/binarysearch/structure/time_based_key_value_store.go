package structure

/*
Problem:
Ref: https://leetcode.com/problems/time-based-key-value-store/

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap()
- Initializes the object of the data structure.
void set(String key, String value, int timestamp)
- Stores the key with the value at the given time timestamp.
String get(String key, int timestamp)
- Returns a value such that set was called previously, with timestamp_prev <= timestamp.
- If there are multiple such values, it returns the value associated with the largest timestamp_prev.
- If there are no values, it returns ""

Constraints:
- key, value length; [1, 100]
- key, value consist lowercase English letters and digits
- timestamp: [1, 10^7]
- At most 2 * 105 calls will be made to set and get.
- All the timestamps of set are strictly increasing.
*/

type TimeBasedValue struct {
	timestamp int
	value     string
}

type TimeMap struct {
	Data map[string][]TimeBasedValue
}

func Constructor() TimeMap {
	return TimeMap{
		Data: make(map[string][]TimeBasedValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Data[key] = append(this.Data[key], TimeBasedValue{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, exists := this.Data[key]
	if !exists {
		return ""
	}

	l, r := 0, len(values)-1
	for l <= r {
		m := (l + r) / 2
		if values[m].timestamp < timestamp {
			l = m + 1
		} else if values[m].timestamp > timestamp {
			r = m - 1
		} else {
			return values[m].value
		}
	}

	if l == 0 {
		return ""
	} else {
		return values[l-1].value
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
