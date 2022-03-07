package pdf

const TPL = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <link rel="shortcut icon" href="./favicon.ico" />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      href="https://fonts.googleapis.com/css2?family=Nunito:wght@200;400;600;800;900&display=swap"
      rel="stylesheet"
    />

    <title>HTML 5 Boilerplate</title>
    <style>
      body,
      html {
        padding: 0;
        margin: 0;
        border-color: #ffffff;
        font-size: 20px;
        font-family: "Nunito", sans-serif;
      }

      .title {
        margin: 0;
        display: flex;
        justify-content: left;
        align-items: center;
      }

      h1 {
        color: #a42d2d;
        margin-left: 100px;
      }

      .top {
        position: relative;
        margin: 0 50px;
      }

      .logo {
        width: 200px;
        height: 140px;
      }

      .content-container {
        padding: 50px 100px;
      }

      .table {
        width: 100%;
        padding: 50px 0;
      }

      .balances {
        margin-top: 150px;
      }

      .sender-info {
        margin: 30px 0 50px 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }

      .sender-info-details,
      .sender-info-dates {
        width: 40%;
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        margin: 0;
        padding: 0;
      }

      .sender-info-details-1,
      .sender-info-details-2 {
        display: flex;
        align-items: center;
      }

      .name {
        font-weight: bold !important;
        width: 50%;
        text-align: left;
        margin: 0;
      }

      .info {
        margin: 0;
        width: 50%;
        text-align: left;
      }
    </style>
  </head>

  <body>
    <div class="content-container">
      <div class="top">
        <div class="title">
          <img class="logo" src="./logo.png" />

          <h1>Statements History</h1>
        </div>
        <div class="balances">
          <div class="sender-info-details">
            <div class="sender-info-details-1">
              <div class="name">Opening Balance</div>
              <div class="info">45453.32</div>
            </div>
            <div class="sender-info-details-2">
              <div class="name">Closing Balance</div>
              <div class="info">56789.00</div>
            </div>
          </div>
        </div>
        <div class="sender-info">
          <div class="sender-info-details">
            <div class="sender-info-details-1">
              <div class="name">Account Number</div>
              <div class="info">01123456789</div>
            </div>
            <div class="sender-info-details-2">
              <div class="name">Account Name</div>
              <div class="info">Maluki Muthusi</div>
            </div>
          </div>

          <div class="sender-info-dates">
            <div class="sender-info-details-1">
              <div class="name">Date From</div>
              <div class="info">23/01/2022</div>
            </div>
            <div class="sender-info-details-2">
              <div class="name">Date To</div>
              <div class="info">04/03/2022</div>
            </div>
          </div>
        </div>
        <hr />
      </div>

      <div>
        <table class="table">
          <tr>
            <th>Date</th>
            <th>Account Name</th>
            <th>Account Number</th>
            <th>Credit/Debit</th>
            <th>Amount</th>
            <th>Description</th>
          </tr>
        </table>
      </div>
    </div>
  </body>
</html>
`
