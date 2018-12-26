<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Bootswatch: Default</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
  </head>
<body>
<!-- 横並びに2要素並べる -->
<div class="row">
  <div class="col-md-6">
    <div class="container"><!-- [start] container for password information on left side -->
      <div class="card"><!--card for  password information-->
        <div class="card-body">
          <h4 class="card-title">Password information</h4>
          <div class="form-group">
            <label for="inputTitle">Title</label>
            <input type="text" class="form-control" id="inputTitle" aria-describedby="title" value="Enter title" readonly><!--readonlyかどうかの要素を切り替える。登録済み設定を選択されたらなんやかんや-->
          </div>
          <div class="form-group">
            <label for="inputPassword">Password</label>
            <input type="password" class="form-control" id="inputPassword" placeholder="Password">
          </div>
          <div class="form-group">
            <label for="algorithm">Algorithm select</label>
            <select class="form-control" id="algorithmSelect">
              {{range $index, $algorithm := .Algorithms}}
              <option>{{$algorithm}}</option>
              {{end}}
            </select>
          </div>
          <div class="form-group">
            <label for="extraInfo">extra Info</label>
            <input type="text" class="form-control" id="extraInfo" aria-describedby="extraInfo to get new password string" placeholder="Enter extra info if you want">
          </div>
          <button type="submit" class="btn btn-primary">Generate password</button>
          <button type="submit" class="btn btn-primary">Save setting</button>
        </div><!--class="card-body-->
      </div><!--<div class="card"> card for  password information-->
      <div class="card"><!-- card for password result-->
        <div class="card-body">
          <h4 class="card-title">Generated password</h4>
          <input type="text" class="form-control form-control-lg" id="inputLarge" value="" placeholder=".form-control-lg" readonly>
        </div>
      </div><!--<div class="card"> card for password result-->
    </div><!-- [end] container for password information on left side -->
  </div>
  <div class="col-md-4">
    <div class="container"><!-- [start] container for user information right side -->
      <div class="card"><!--card for login information-->
        <h4 class="card-title">Login</h4>
        <div class="card-body">
          <div class="form-group">
            <label for="username">Username</label>
            <input type="text" class="form-control" id="Username" placeholder="Enter username">
          </div>
          <div class="form-group">
            <label for="password">Login password</label>
            <input type="password" class="form-control" id="LoginPassword" placeholder="Enter username">
          </div>
        </div>
      <div><!--<div class="card"> card for  password information-->
      <div class="card"><!--card for table information-->
        <table class="table table-striped">
          <thead>
            <th scope="col">Password Settings</th>
          </thead>
          <tbody id=user_password_settings>
            <!-- ここも設定一覧で収める。ただしユーザー毎に設定値が異なるので、jsで取得させる-->
            <!--
            <tr><td>Column content</td></tr>
            <tr><td>Column content</td></tr>
            <tr><td>Column content</td></tr>
            <tr><td>Column content</td></tr>
            <tr><td>Column content</td></tr>
            <tr><td>Column content</td></tr>
            -->
          </tbody>
        </table>
      </div>
    </div><!-- [end] container for user information right side -->
  </div>
</div>
</body>
