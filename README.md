# A-December-of-Algorithms-2024

<div align="left">
<h1>
    <p align="center"><img alt="header" src="https://github.com/user-attachments/assets/da6f80f2-06db-4578-9833-f635b85e3da6" width="500"></img></p>

</h1>
Welcome to A December of Algorithms 2024.

Building on the success and enthusiasm of previous years, we’re excited to bring you a fresh selection of algorithms to explore and implement this December!

Each Day, Each Algorithm ;)
Finish them all to get a certificate :)

**Send a pull request only after completing all 31 algorithms.**

**Please submit all PRs on or before January 21st 11:59 PM IST.**

## What Do I Do?

We have a small collection of algorithms, one for every day of the month. Scroll down to take a look at them.

All you need to do is fork this repository, implement all 31 algorithms and send a pull request over to us.

Check out our FAQ for more information.

## Index

- [**December 01 - The Vanishing Number**](#december-01---the-vanishing-number)
- [**December 02 - The Wave Sort Challenge**](#december-02---the-wave-sort-challenge)
- [**December 03 - Alternating Square Arrangement**](#december-03---alternating-square-arrangement)
- [**December 04 - Plant Growth Tracker**](#december-04---plant-growth-tracker)
- [**December 05 - Josephus Problem**](#december-05---josephus-problem)
- [**December 06 - Target Pair Finder**](#december-06---target-pair-finder)
- [**December 07 - The Magical Tower**](#december-07---the-magical-tower)
- [**December 08 - Digit Manipulation**](#december-08---digit-manipulation)
- [**December 09 - Customer Return Frequency**](#december-09---customer-return-frequency)
- [**December 10 - Concurrent Task Execution**](#december-10---concurrent-task-execution)
- [**December 11 - The Robot Returns**](#december-11---the-robot-returns)
- [**December 12 - Smart Ticketing System**](#december-12---smart-ticketing-system)
- [**December 13 - Minimum Swap Sorting Problem**](#december-13---minimum-swap-sorting-problem)
- [**December 14 - Split the Squad**](#december-14---split-the-squad)
- [**December 15 -  Holiday Gift Arrangement**](#december-15---holiday-gift-arrangement)




- [**FAQ**](#faq)

## Algorithms

### December 01 - The Vanishing Number

#### Problem Statement

```
Imagine a race event with N participants, each having a unique bib number from 1 to N.
After the race, the event organizers collect the bib numbers of all participants who showed
up. However, one participant didn’t show up, and their bib number is missing. The task is to
find out which bib number is missing in order to identify the participant who was absent.
```

```
Given an integer N and an array of size N-1 containing N-1 numbers between 1 to N. Find
the number(between 1 to N), that is not present in the given array.
```

![image](https://github.com/user-attachments/assets/177fbe2f-9248-481b-adb6-61dfcea59582)

Example 1:

```
Input Format:
N = 5, array[] = {1,2,4,5}
Result: 3
```

```
Explanation:
In the given array, number 3 is missing. So, 3 is the answer.
```

Example 2:

```
Input Format:
N = 3, array[] = {1,3}
Result: 2
```

### December 02 - The Wave Sort Challenge

#### Problem Statement

```
Imagine you're at a music concert, surrounded by thousands of fans.
As the music plays, the crowd begins to move in a rhythmic pattern, rising and falling like a wave.
The energy is contagious, and everyone follows the same wave-like motion.
Now, think of the crowd as an unsorted array of integers, and your job is to arrange them in such a way that
the heights of the fans rise and fall, just like the waves at the concert.
Your task is to organize the array in a similar wave-like fashion!
```

```
Sort an array in wave form
Given an unsorted array of integers, sort the array into a wave array. An array arr[0..n-1] is sorted in wave form if:
arr[0] >= arr[1] <= arr[2] >= arr[3] <= arr[4] >= ….

```

![image1](https://github.com/user-attachments/assets/7c46bc27-91f9-4a89-8731-37efa558a939)

Example 1:

```
Input:
arr[] = {10, 5, 6, 3, 2, 20, 100, 80}
Output:
arr[] = {10, 5, 6, 2, 20, 3, 100, 80}

```

```
Explanation:
Here you can see {10, 5, 6, 2, 20, 3, 100, 80}. The first element is larger than the second, and the same pattern repeats again.
Large element – small element – large element – small element, and so on.
It can also be in the opposite pattern (small element – large element – small element – large element).
All you need to maintain is the up-down fashion, which represents a wave. There can be multiple answers.

```

Example 2:

```
Additional Scenario
Input:
arr[] = {1, 2, 3, 4, 5, 6}
Output:
arr[] = {2, 1, 4, 3, 6, 5}

```

```
Explanation:
In this case, we swap adjacent elements to form a wave-like array.
{2, 1, 4, 3, 6, 5} satisfies the condition where elements are alternately large and small.

```

### December 03 - Alternating Square Arrangement

#### Problem Statement

```
You are given integers ￼ (number of red squares) and ￼ (number of blue squares).
Your task is to arrange the squares in a sequence such that:
No two adjacent squares have the same color.
If it is not possible to create such an arrangement, return "Not possible".

```

```
Input

Two integers:
R￼: Number of red squares.
B￼: Number of blue squares.

Output

A valid arrangement as a string (e.g., "RBRBR" or "BRBRB") that satisfies the condition, or "Not possible" if no valid arrangement exists.


```

Sample Input and Output:
Example 1:

```
Input:
R = 3
B = 2

Output:
"RBRBR"
```

Explanation:

```
The sequence "RBRBR" satisfies the condition since no two adjacent squares are the same color.
```

Example 2:

```
Input:
R = 4
B = 2

Output:
"Not possible"
```

Explanation:

```
It is impossible to arrange 4 red squares and 2 blue squares without
having two adjacent squares of the same color. Hence, the output is "Not possible".

```

### December 04 - Plant Growth Tracker

#### Problem Statement

```
You are designing a program to help a gardener plan how to grow a special type of plant. The plants
grow in an interesting pattern:
• It takes 2 months for the plant to before it starts increasing in numbers.
• Starting from the third month, the number of plants in a month is the sum of the plants in the
previous two months.
For example, in the third month, there are 2 plants (1 + 1), in the fourth month, there are 3 plants (1 +
2), and so on.
```

![image](https://github.com/user-attachments/assets/4810c82d-5cc3-4ab8-a247-64d8d8133e44)

```
However, the gardener needs your help to write a program that predicts the number of plants for
any given number of months, n. Here's the catch:
• The gardener's old computer struggles to handle repeated function calls. As a result, you
cannot use recursion to solve this problem.
```

Example 1:

```
Input: No of Months: 10
Output: 55
```

Explanation:

```
We apply Fibonacci sequence here. Since the growth of the plant perfectly mirrors with Fibonacci
algorithm, as the gardener's computer is old, we have to use for loops and if statements to generate
the sequence.
```

Example 2:

```
Input: No of Months: 12
Output: 144
```

### December 05 - Josephus Problem

#### Problem Statement

```
A total of n people are standing in a circle, and you are one of them playing a game.
Starting from a person, k persons will be counted in order along the circle, and the kth person will be killed.
Then the next k persons will be counted along the circle, and again the kth person will be killed. This cycle will
continue until only a single person is left in the circle.
If there are 5 people in the circle in the order A, B, C, D, and E, you will choose k=3. Starting from A, you will
count A, B and C. C will be the 3rd person and will be killed. Then you will continue counting from D, E and
then A. A will be third person and will be killed.
```

![image](https://github.com/user-attachments/assets/5925562d-fc20-48b7-9114-016a0f2819e0)

```
You will be given an array where the first element is the first person from whom the counting will start and
the subsequent order of the people. You want to be the last one standing. Determine the index at which you
should stand to survive the game. Return an integer denoting safe position.
```

Example 1:

```
Input: n = 3, k = 2
Output: 3
```

Explanation:

```
There are 3 persons so skipping 1 person i.e 1st person 2nd person will be killed. Thus the safe
position is 3.
```

Example 2:

```
Input: n = 5, k = 3
Output: 4
```

Explanation:

```
There are 5 persons so skipping 2 person i.e 3rd person will be killed. Thus the safe position is 4.
```
### December 06 - Target Pair Finder

#### Problem Statement
```
From the initialised list of integers and a target sum by the user,
find all unique pairs of numbers from the list that add up to the target.
You can use nested loops and conditionals.
```


Sample I/O 1:
```
INPUT: 1) A list of integers: numbers = [2, 4, 3, 7, 1, 5].
              2) A target sum: target = 6.
OUTPUT:  Unique pairs are [(2, 4), (1, 5)]
```
Sample I/O 2:
```
INPUT:  1) A list of integers: numbers = [10, 15, 3, 7, 8, 12, 5].
              2)  A target sum: target = 20.
OUTPUT: Unique pairs are [(10, 10), (8, 12), (15, 5)]
Explanation:
The Target Pair Finder problem is about finding pairs of numbers in a list that add up to a specific target value. From the above example we can see that (10, 10): The first 10 and the second 10 in the list add up to 20, so thereby find all sets of unique pair from the given list summing up to the target value.
```
### December 07 - The Magical Tower

#### Problem Statement
```
You are tasked with designing a Magical Tower for a kingdom. The tower has multiple floors, and each floor is supported by a triangular arrangement of magical stones called the Pascal Stones. These stones have unique properties:

The stones at the edges of the triangle are always marked as 1.
The inner stones on each floor are infused with power equal to the sum of the two stones directly above them from the previous floor.
The first floor of the tower contains only a single stone ([1]), and subsequent floors are built according to the rules above. The kingdom's wizard has asked you to construct the first N floors of the Magical Tower.

Your task: Write a program that generates the arrangement of stones for the first N floors of the tower.
```

![PascalTriangleAnimated2](https://github.com/user-attachments/assets/c7be7eab-bd66-4947-93be-e7874f2670d3)

Sample I/O 1:
```
INPUT: numRows=5
OUTPUT:  [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```
Explanation:
```
The first 5 floors of the Magical Tower are constructed as follows:
Floor 1: [1]
Floor 2: [1, 1]
Floor 3: [1, 2, 1] → The second stone (2) is created by adding the two stones directly above it (1 + 1).
Floor 4: [1, 3, 3, 1] → The second stone (3) is created by adding 1 + 2, and the third stone (3) is created by adding 2 + 1.
Floor 5: [1, 4, 6, 4, 1] → The second stone (4) is created by adding 1 + 3, the third stone (6) by adding 3 + 3, and the fourth stone (4) by adding 3 + 1.
Each floor is constructed based on the magical rule that the value of each inner stone is the sum of the two stones directly above it.
```
Sample I/O 2:
```
INPUT:  numRows=1
OUTPUT: [1]
```
### December 08 - Digit Manipulation

#### Problem Statement
```
Write a program to calculate the Digit Square Sum for all numbers from 1 to a given positive integer
N.
The Digit Square Sum of a number is calculated by squaring each digit of the number and then
summing up the squares.
```
```
Your task is to write a function that:
1. Takes an integer N as input.
2. Computes the Digit Square Sum for each number from 1 to N.
3. Returns the total sum of these values.
```

Example 1:
```
For N = 12, the program calculates the following:
- 1 -> 1^2 = 1
- 2 -> 2^2 = 4
- 3 -> 3^2 = 9
- ...
- 12 -> 1^2 + 2^2 = 1 + 4 = 5
```
```
The total sum is 1 + 4 + 9 + 16 + 25 + 36 + 49 + 64 + 81 + 1 + 2 + 5 = 293.
Output:
The program should return the total Digit Square Sum.
```


### December 09 - Customer Return Frequency

#### Problem Statement

```
You are managing an e-commerce platform and you have a list of customer return
frequencies (how many times a customer has returned products). Your task is to find the total number of
customers who have returned products exactly once.
```


Example 1:

```
Input Format:
returns = [2, 1, 5, 1, 0, 3, 1, 4, 1]
Result: 4
```

```
Explanation:
The list returns = [2, 1, 5, 1, 0, 3, 1, 4, 1] represents the return frequency of each customer.
 We are looking for customers who have returned products exactly once, so we need to count how
many times 1 appears in the list.
```

Example 2:

```
Input Format:
returns = [4, 3, 7, 2, 1, 0, 2, 1, 3]
Result: 2
```
### December 10 - Concurrent Task Execution
#### Problem Statement
```
You are given a list of tasks, where each task has a unique identifier and a list of dependencies (other
tasks that must be completed before this task can be executed). Your goal is to determine a valid
order to execute the tasks using concurrency wherever possible.
You must account for the dependencies and ensure no task runs before its dependencies are
completed. If no valid execution order exists (i.e., there is a cyclic dependency), return an error
message.
```
Input:
```
• A list of tasks, where each task is represented as a pair (task_id, dependencies).
o task_id is a unique identifier for the task (e.g., an integer or string).
o dependencies is a list of task IDs that must be completed before the given task can
run.
```
Output:
```
If a valid task execution order exists, return a list of lists, where each sublist contains the task
IDs that can be executed concurrently at that step.
• If no valid order exists (i.e., a circular dependency is found), return the message "Error: Cyclic
dependency detected".
```

![image](https://github.com/user-attachments/assets/e848e914-4128-457a-be5b-6f64699b86df)

Sample 1:
```
Input:
tasks = [
 ("A", []),
 ("B", ["A"]),
 ("C", ["A"]),
 ("D", ["B", "C"]),
 ("E", ["D"])
]
```
```
Output:
[["A"], ["B", "C"], ["D"], ["E"]]
```
Explanation:
```
• "A" has no dependencies, so it runs first.
• "B" and "C" both depend only on "A", so they can run concurrently.
• "D" depends on both "B" and "C", so it runs after them.
• "E" depends on "D", so it runs last
```
Sample 2:
```
Input:
tasks = [
 ("A", ["B"]),
 ("B", ["A"])
]
```
```
Output:
"Error: Cyclic dependency detected"
```

Explanation:
```
• "A" depends on "B" and "B" depends on "A", creating a cycle.
```

### December 11 - The Robot Returns
#### Problem Statement
```
There is a robot starting at the position (0, 0), the origin, on a 2D plane. Given a
sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.
You are given a string moves that represents the move sequence of the robot where moves[i] represents
its i
th move. Valid moves are 'R' (right), 'L' (left), 'U' (up), and 'D' (down).
Return true if the robot returns to the origin after it finishes all of its moves, or false otherwise.
```
```
Note: The way that the robot is "facing" is irrelevant. 'R' will always make the robot move to the right
once, 'L' will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the
same for each move.
```
![image](https://github.com/user-attachments/assets/d6399796-727c-417b-9f6c-c68a4bc21743)

Sample 1:
```
Input: moves = "UD"
Output: true
Explanation: The robot moves up once, and then down once. All moves have the same magnitude, so it
ended up at the origin where it started. Therefore, we return true.
```
Sample 2:
```
Input: moves = "LL"
Output: false
Explanation: The robot moves left twice. It ends up two "moves" to the left of the origin. We return false
because it is not at the origin at the end of its moves.
```
Reference: For more information on topological sorting and dependency resolution, check out this guide on https://www.geeksforgeeks.org/topological-sorting/


### December 12 - Smart Ticketing System
#### Problem Statement
```
You are tasked with designing a Smart Ticketing System for a popular concert. The system
manages ticket requests using a queue data structure but with additional complexity:
1. Priority Handling:
Some customers are marked as VIPs (designated by a VIP tag in their request). VIP
customers have higher priority and are served before regular customers, regardless of
their position in the queue. However, among VIPs or regular customers, the requests are
handled in the order they are received (FIFO).
2. Dynamic Ticket Allocation:
Each request includes the number of tickets the customer wants. If the requested tickets
exceed the remaining tickets, the system will allocate all remaining tickets to the
customer.
3. Queue Operation:
If a customer receives fewer tickets than requested due to limited availability, the request
is still considered processed, and the next customer in the queue is served.
You must implement a program that processes these ticket requests and returns the result of
each transaction.
```
Contraints:
```
1. 2. 3. The system starts with N tickets available.
Each request is represented as a string in the format "CustomerName
NumberOfTickets [VIP]"
If [VIP] is not present, the customer is treated as a regular customer.
Requests are processed until all tickets are sold out or the queue is empty
```
Sample 1:
```
Input:
N = 5
requests = ["John 2 VIP" ,"Alice 3", "Bob 2" , "Charlie 1 VIP"]
Output:
["John purchased 2 tickets", "Charlie purchased 1 tickets"
, "Alice purchased 2 tickets", Bob was not served"]
Explanation:
"John 2 VIP" is served first because he is a VIP.
"Charlie 1 VIP" is served next, as he is also a VIP.
"Alice 3" is served, but only 2 tickets are left, so she gets 2.
"Bob 2" cannot be served as there are no tickets remaining.
```
Sample 2:
```
Input:
N = 10
requests = ["Eve 4","Diana 3 VIP","Adam 5","Frank 6 VIP"]
Output:
["Diana purchased 3 tickets","Frank purchased 6 tickets","Eve purchased tickets", "Adam was not served"]
```
### December 13 - Minimum Swap Sorting Problem
#### Problem Statement
```
John has a list of unique integers that he wants to sort in ascending order.
However, he can only sort the list by swapping two elements at a time.
The "cost" of each swap is 1 unit.
Your task is to determine the minimum cost
(i.e., the minimum number of swaps required) to sort the list.
```
Example 1:
```
Sample Input 1:
5
4 3 1 2 5
Sample Output 1:
3
```
Explanation:
```
The given list is [4, 3, 1, 2, 5].
Swap 4 and 1: [1, 3, 4, 2, 5]
Swap 3 and 2: [1, 2, 4, 3, 5]
Swap 4 and 3: [1, 2, 3, 4, 5]
Total swaps = 3. Hence, the minimum cost is 3.
```
Example 2:
```
Sample Input 2:
4
2 3 4 1
Sample Output 2:
3
```
```
Input Format:
The first line contains an integer, N, the total number of integers in the list.
The second line contains N space-separated integers representing the list.
Output Format:
An integer representing the minimum cost (number of swaps) required to sort the list.
```
References: This problem is inspired by sorting algorithms and cycle detection in graphs.

### December 14 - Split the Squad
#### Problem Statement
```
Alice has N students in his class, numbered 1 through N. Each student has
expertise in a subject numbered Ai

. Alice has to divide the students into two

teams Team 1 and Team 2, such that:
1. Each student belongs to exactly one team.

2. The uniqueness of each team is exactly K.

3. Additionally, the difference in the number of students between the two teams
must not exceed D.
```
```
The uniqueness of a team is defined as the number of distinct subjects such that
there is at least one student in the team with expertise in the subject. For
example, the uniqueness of a team denoted by A = [1, 3, 1, 4, 4] is 3.
Alice wants to know if it is possible to distribute the students into two teams
satisfying the conditions.
```
```
Input format
The first line contains an integer T, the number of test cases.
For each test case:
The first line contains three integers N, K, and D.
The second line contains N integers A1
,A2
,....,An
```
```
Output format
For each test case, print YES if Alice is able to make two teams satisfying the
given conditions, otherwise print NO.
```
```
Constraints

1 ≤ T ≤ 105
2 ≤N ≤ 105
1 ≤ K ≤ N
1 ≤ D ≤ N
1 ≤ Ai ≤ N
```
![image](https://github.com/user-attachments/assets/3111d8a4-a02d-4804-b7e4-f2257d631abe)

```
Explanation
Test Case 1:
Divide students into Team 1 = [1, 2, 2] and Team 2 = [3, 4, 4]. Both teams have
a uniqueness of 2, and the difference in the number of students is 0 (≤ 2).
Output is YES.

Test Case 2:
No way to divide the students into two teams with both having a uniqueness of
3 while keeping the size difference ≤ 1.
Output is NO.
```

References
https://www.geeksforgeeks.org/greedy-algorithms
https://www.geeksforgeeks.org/hashing-data-structure



### December 15 - Holiday Gift Arrangement
#### Problem Statement
```
It’s December, and Santa is preparing to deliver gifts. He has N houses to visit,
each requiring a certain number of gifts. Santa’s sleigh can carry a maximum of W gifts at a time.
You are given:
	•	An array houses[] where each element represents the number of gifts required at a house.
	•	The maximum carrying capacity W of Santa’s sleigh.
Santa needs to minimize the number of trips required to deliver all the gifts.
Each trip can serve one or more consecutive houses as long as the total number of gifts does not exceed W.
Write a function minTrips(houses, W) that returns the minimum number of trips Santa needs to deliver the gifts
Here is an artistic depiction of Santa Claus preparing for his December gift deliveries.
It captures the festive and cheerful atmosphere with snow-covered houses and a sleigh loaded with gifts.
```
Example:
```
Input:
houses = [2, 3, 5, 2, 1]
W = 6
Output:
3
Explanation:
	•	Trip 1: Deliver to house 1 and 2 (2 + 3 gifts = 5 ≤ 6).
	•	Trip 2: Deliver to house 3 (5 gifts = 5 ≤ 6).
	•	Trip 3: Deliver to house 4 and 5 (2 + 1 gifts = 3 ≤ 6).

```
Hints:
```
Use a greedy approach to group consecutive houses.
```
# FAQ

#### Who can join the Challenge?

Anyone who is passionate about coding and can dedicate a little time a day for the challenge for the next 31 days.

#### When should I submit the pull request?

You don't need to submit it everyday. Just submit it once you're done with all 31 algorithms.

#### What if I'm not able to code every day?

Not a problem. While coding every day is nice, we understand that other commitments might interfere with it.

Plus its holiday season. So you don't have to solve one problem every day.

Go at your own pace. One per day or 7 a week or even all 30 in a day.

#### What language should I use to code?

Anything! New to GoLang? Best way to practice it.

Wanna find out what all this hype about Python is? Use it!

Any and all languages are welcome.

Maybe you could try using a different language for every problem as a mini-challenge?

#### Fork? Pull request? What is all that? I don't know how to use GitHub!

If you are new to Git or GitHub, check out this out [GitHub](https://guides.github.com/activities/hello-world/)

#### Where are the rest of the problems?

Our code ninjas are hard at work preparing the rest of the problems. Don't worry, they'll be up soon.

#### How should I complete these programs?

We have a folder for each day of the month. Simply complete your code and move the file into that folder.

Be sure to rename your file to the following format: `language_username` or `language_username_problemname`
Some examples:
`python3_exampleUser.py`
`c_exampleUser.c`

**Please do not modify any existing files in the repository.**

#### I forked the repository but some problems were added only after that. How do I access those problems?

Not to worry! Open your nearest terminal or command prompt and navigate over to your forked repository.

Enter these commands:

```bash
git remote add upstream https://github.com/SVCE-ACM/A-December-of-Algorithms-2024.git
git fetch upstream
git merge upstream/main
```

If you're curious, the commands simply add a new remote called upstream that is linked to this repository. Then it 'fetches' or retrieves the contents of the repository and attempts to merge it with your progress.
Note that if you've already added the upstream repository, you don't need to re-add it in the future while fetching the newer questions.

#### I received a merge error. What do I do?

This shouldn't happen unless you modify an existing file in the repository. There's a lot of potential troubleshooting that might be needed, but the simplest thing to do is to make a copy of your code outside the repository and then clone it once again. Now repeat the steps from the answer above. Merge it and then add your code. Now proceed as usual. :)

#### I'm facing difficulties with/need help understanding a particular question.

Open up an [issue](https://github.com/SVCE-ACM/A-December-of-Algorithms-2021/issues) on this repository and we'll do our best to help you out.

###### [[Back to Top]](#----)

![wave](http://cdn.thekrishna.in/img/common/border.png)

```

```
