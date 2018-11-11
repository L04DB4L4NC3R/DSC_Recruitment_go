## DSC Recruitment portal
Record and show recruitment details

<br />

| route |  type  |  data  |  response  |
|---|---|---|---|
| /record | POST | User object | New record added | 
| /show | GET with Authorization header | ?reg={REGISTRATION_NUMBER} (optional) | []User object or User object |
| /show/{applicantType} | GET with Authorization header |  applicantType="technical" etc.  |  []User object |
| /manager/record | POST | Management object | OK|
| /manager/show | GET | ?reg={REGISTRATION_NUMBER} (optional) | []Management object or Management object|


<br />
<br />

### User object

<br />

```json
{
    "Name":"dhruv sharma",
    "Email":"dhruvsharma1016@gmail.com",
    "Reg":"16BCE0955",
    "ApplicantType":"design"
}
```

<br />
<br />


### Management object

<br />

```json
{
    "reg":"16BCE0955",
    "q1":"answer",
    "q2":"answer",
    "q3":"answer",
    "q4":"answer",
    "q5":"answer",
    "q6":"answer",
    "q7":"answer",
    "q8":"answer",
    "q9":"answer",
    "q10":"answer"
}
```