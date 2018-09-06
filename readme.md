# Kata Bowling
Create a program, which, given a valid sequence of rolls for one line of American Ten-Pin Bowling, produces the total score for the game. Here are some things that the program will not do:

We will not check for valid rolls.
We will not check for correct number of rolls and frames.
We will not provide scores for intermediate frames.
Depending on the application, this might or might not be a valid way to define a complete story, but we do it here for purposes of keeping the kata light. I think you’ll see that improvements like those above would go in readily if they were needed for real.

We can briefly summarize the scoring for this form of bowling:

Each game, or “line” of bowling, includes ten turns, or “frames” for the bowler.
In each frame, the bowler gets up to two tries to knock down all the pins.
If in two tries, he fails to knock them all down, his score for that frame is the total number of pins knocked down in his two tries.
If in two tries he knocks them all down, this is called a “spare” and his score for the frame is ten plus the number of pins knocked down on his next throw (in his next turn).
If on his first try in the frame he knocks down all the pins, this is called a “strike”. His turn is over, and his score for the frame is ten plus the simple total of the pins knocked down in his next two rolls.
If he gets a spare or strike in the last (tenth) frame, the bowler gets to throw one or two more bonus balls, respectively. These bonus throws are taken as part of the same turn. If the bonus throws knock down all the pins, the process does not repeat: the bonus throws are only used to calculate the score of the final frame.
The game score is the total of all frame scores.

- X X X X X X X X X X X X (12 rolls: 12 strikes) = 10 frames * 30 points = 300
- 9- 9- 9- 9- 9- 9- 9- 9- 9- 9- (20 rolls: 10 pairs of 9 and miss) = 10 frames * 9 points = 90
- 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/5 (21 rolls: 10 pairs of 5 and spare, with a final 5) = 10 frames * 15 points = 150

# Function And Parameter Naming Conventions

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
api.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_RequestToOtherWebservice_Input_ApplicationForm_Should_Be_ApplicationResult
```

## Variable Name
- ชื่อตัวแปรเป็น **camelCase** เช่น
```
policy, applicationForms, applicationResult, policyResult
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
insurers
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
ResultDay, ResultMonth
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวพิมพ์เล็กก่อน เว้นแต่เมื่อมีการใช้ข้าม package ถึงจะใช้ Capital Case เช่น
```
Hour, Minute, url
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน

## คำสั่ง Run Test
### ค่าสั่ง Run Acceptance Test (Robot)

```
robot duration.robot
```

### คำสั่ง Run Acceptance Test (API)
```
newman run filename
```

## Additional
[See more git and go command](https://github.com/ImKK-000/git-and-go-step)