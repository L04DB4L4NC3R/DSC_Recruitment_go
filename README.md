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
    "Reg":"16BCE0955",
    "Q1":"answer",
    "Q2":"answer",
    "Q3":"answer",
    "Q4":"answer"
}
```