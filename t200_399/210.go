package t200_399

var ans []int
var froms map[int][]int
var vis []bool

//class Solution {
//public int[] findOrder(int numCourses, int[][] prerequisites) {
//if (numCourses <= 0) {
//return new int[0];
//}
//
//HashSet<Integer>[] adj = new HashSet[numCourses];
//for (int i = 0; i < numCourses; i++) {
//adj[i] = new HashSet<>();
//}
//
//// [1,0] 0 -> 1
//int[] inDegree = new int[numCourses];
//for (int[] p : prerequisites) {
//adj[p[1]].add(p[0]);
//inDegree[p[0]]++;
//}
//
//Queue<Integer> queue = new LinkedList<>();
//for (int i = 0; i < numCourses; i++) {
//if (inDegree[i] == 0) {
//queue.offer(i);
//}
//}
//
//int[] res = new int[numCourses];
//// 当前结果集列表里的元素个数，正好可以作为下标
//int count = 0;
//
//while (!queue.isEmpty()) {
//// 当前入度为 0 的结点
//Integer head = queue.poll();
//res[count] = head;
//count++;
//
//Set<Integer> successors = adj[head];
//for (Integer nextCourse : successors) {
//inDegree[nextCourse]--;
//// 马上检测该结点的入度是否为 0，如果为 0，马上加入队列
//if (inDegree[nextCourse] == 0) {
//queue.offer(nextCourse);
//}
//}
//}
//
//// 如果结果集中的数量不等于结点的数量，就不能完成课程任务，这一点是拓扑排序的结论
//if (count == numCourses) {
//return res;
//}
//return new int[0];
//}
//}
func findOrder(numCourses int, prerequisites [][]int) []int {
	vis = make([]bool, numCourses)
	froms = map[int][]int{}
	ans = []int{}
	rus := make([]int, numCourses)
	for _, v := range prerequisites {
		if _, ok := froms[v[0]]; ok {
			froms[v[0]] = append(froms[v[0]], v[1])
		} else {
			froms[v[0]] = []int{v[1]}
		}
		rus[v[0]]++
	}
	for {
		f := true
		for i := 0; i < numCourses; i++ {
			if rus[i] == 0 && !vis[i] {
				ans = append(ans, i)
				vis[i] = true
				f = false
				for k, v := range froms {
					for _, vv := range v {
						if vv == i {
							rus[k]--
						}
					}
				}
			}
		}
		if len(ans) == numCourses {
			return ans
		}
		if f {
			return []int{}
		}
	}

	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}
