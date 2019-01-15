<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <!--<script type="text/javascript" src="bind.js"></script>-->
    <title>Generate password</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
  </head>
<body>
<!-- 横並びに2要素並べる -->
<nav class="navbar navbar-light bg-light">
  <a class="navbar-brand" href="https://github.com/developer-kikikaikai/CreatingHashPassword">
    <img src="images/logo.png" height="60" class="d-inline-block align-top" alt="">
    Generate password
  </a>
</nav>
<div class="row">
  <div class="col-md-6">
    <div class="container"><!-- [start] container for passphrase information on left side -->
      <div class="card"><!--card for  passphrase information-->
        <div class="card-body">
          <h4 class="card-title">Passphrase information</h4>
          <div class="form-group">
            <label for="inputTitle">Title</label>
            <input type="text" class="form-control" id="inputTitle" aria-describedby="title" value="Enter title" readonly><!--readonlyかどうかの要素を切り替える。登録済み設定を選択されたらなんやかんや-->
          </div>
          <div class="form-group">
            <label for="inputKeyphrase">Keyphrase</label>
            <input type="password" class="form-control" id="inputPassphrase" placeholder="Keyphrase">
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
            <input type="text" class="form-control" id="extraInfo" aria-describedby="extraInfo to get new passphrase string" placeholder="Enter extra info if you want">
          </div>
          <button type="submit" class="btn btn-primary">Generate passphrase</button>
          <button type="submit" class="btn btn-primary">Save setting</button>
        </div><!--class="card-body-->
      </div><!--<div class="card"> card for  passphrase information-->
      <div class="card"><!-- card for passphrase result-->
        <div class="card-body">
          <h4 class="card-title">Generated passphrase</h4>
          <input type="text" class="form-control form-control-lg" id="inputLarge" value="" readonly>
        </div>
      </div><!--<div class="card"> card for passphrase result-->
    </div><!-- [end] container for passphrase information on left side -->
  </div>
  <div class="col-md-4">
    <div class="container"><!-- [start] container for user information right side -->
      <div class="card"><!--card for creating login user-->
        <h4 class="card-title"> Create user</h4>
        <div class="card-body">
          <div class="form-group">
            <label for="username">Username</label>
            <input type="text" class="form-control" id="Username" placeholder="Enter username">
          </div>
          <div class="form-group">
            <label for="passphrase">Login passphrase</label>
            <input type="password" class="form-control" id="LoginPassphrase" placeholder="Enter username">
          </div>
          <button type="submit" class="btn btn-primary">Create</button>
        </div>
      <div><!--<div class="card"> card for  passphrase information-->
      <div class="card"><!--card for updating login information-->
        <h4 class="card-title"> Update user</h4>
        <div class="card-body">
          <div class="form-group">
            <label for="passphrase">Login passphrase</label>
            <input type="password" class="form-control" id="LoginPassphrase" placeholder="Enter username">
          </div>
          <button type="submit" class="btn btn-primary">Update</button>
        </div>
      <div><!--<div class="card"> card for  passphrase information-->
      <div class="card"><!--card for table information-->
        <h4 class="card-title"> Passphrase Settings</h4>
        <div class="card-body">
          <table class="table table-striped">
            <thead>
              <th scope="col">Passphrase Settings</th>
            </thead>
            <tbody id=user_passphrase_settings>
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
            <button type="submit" class="btn btn-primary">Get setting</button>
          </table>
        </div><!--<div class="card"> card for  passphrase information-->
      </div><!--<div class="card"> card for  passphrase information-->
    </div><!-- [end] container for user information right side -->
  </div>
</div>
</body>
