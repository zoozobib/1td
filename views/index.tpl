<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }


    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
    .description {
      padding-top: 20px;
    }
  </style>
</head>
<!--
Uname:E*** RegTime:2016-05-28 13:28:32 ID: IvestTime:2016-05-29 21:21:31 Expand: IvestMoney:100.00 Mphone:158****1731 AuthName:1 Source:易通贷
-->
<body>
  <header>
    <h1 class="logo">1td的数据</h1>
    <div class="description">
      <tbody>
        <table>
          <tr>
            <td>Uname</td>
            <td>RegTime</td>
            <td>ID</td>
            <td>IvestTime</td>
            <td>Expand</td>
            <td>IvestMoney</td>
            <td>Mphone</td>
            <td>AuthName</td>
            <td>Source</td>
          </tr>
          {{range $index, $elem := .content}}
          <tr>
            <td> {{$elem.Uname}}</td>
            <td> {{$elem.RegTime}}</td>
            <td> {{$elem.ID}}</td>
            <td> {{$elem.RegTime}}</td>
            <td> {{$elem.Expand}}</td>
            <td> {{$elem.IvestMoney}}</td>
            <td> {{$elem.Mphone}}</td>
            <td> {{$elem.AuthName}}</td>
            <td> {{$elem.Source}}</td>
          </tr>
          {{end}}
        </table>
      </tbody>
    </div>
  </header>
  <footer>
    <div class="author">
      Official website: 联系我? ／
      Contact me: webmaster@***.com
    </div>
  </footer>
  <div class="backdrop"></div>
</body>
</html>
