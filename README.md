## DSC Recruitment portal
Record and show recruitment details

<br />

| route |  type  |  data  |  response  |
|---|---|---|---|
| /record | POST | []User object | New record added | 
| /show | GET with Authorization header | ?reg=REGISTRATION_NUMBER (optional) | []User object |
| /show/{applicantType} | GET with Authorization header |  applicantType="technical" etc.  |  []User object |


<br />
<br />

### User object

<br />

```json
[{
    "Name":"dhruv sharma",
    "Email":"dhruvsharma1016@gmail.com",
    "Reg":"16BCE0955",
    "ApplicantType":"design"
}]
```