<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <script type="text/javascript" src="js/main.js"></script>
    <title>Generate passphrase</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
  </head>
<body>
<!-- 横並びに2要素並べる -->
<nav class="navbar navbar-light bg-light">
  <a class="navbar-brand" href="https://github.com/developer-kikikaikai/CreatingHashPassword">
    <img src="images/logo.png" height="60" class="d-inline-block align-top" alt="">
    Generate passphrase
  </a>
</nav>
<div class="row">
  <div class="col-md-6">
    <div class="container"><!-- [start] container for passphrase information on left side -->
      <div class="card"><!-- card for passphrase result-->
        <div class="card-body">
          <h4 class="card-title">Generate passphrase</h4>
          <div class="form-group">
            <label for="inputKeyphrase">From Passphrase setting and follow Keyphrase:</label>
            <input type="password" class="form-control" id="keyphrase_In_PassphraseSetting" placeholder="Keyphrase">
          </div>
          <button type="submit" class="btn btn-primary" id="SubmitGeneratePassphrase">Let's Generate!</button>
          <h4 class="card-title">Generated passphrase</h4>
          <input type="text" class="form-control form-control-lg" id="result_In_Generated_passphrase" value="" readonly>
        </div>
      </div><!--<div class="card"> card for passphrase result-->
      <div class="card"><!--card for  passphrase information-->
        <div class="card-body">
          <h4 class="card-title">Passphrase setting</h4>
          <div class="form-group">
            <label for="inputTitle">Title</label>
            <input type="text" class="form-control" id="title_In_PassphraseSetting" aria-describedby="title" value="Enter title"><!--readonlyかどうかの要素を切り替えるようにしたい。登録済み設定を選択されたらなんやかんや-->
          </div>
          <div class="form-group">
            <label for="algorithm">Algorithm select</label>
            <select class="form-control" id="algorithmSelect_In_PassphraseSetting">
              {{range $index, $algorithm := .Algorithms}}
              <option>{{$algorithm}}</option>
              {{end}}
            </select>
          </div>
          <div class="form-group">
            <label for="extraInfo">extra Info</label>
            <input type="text" class="form-control" id="extraInfo_In_PassphraseSetting" aria-describedby="extraInfo to get new passphrase string" placeholder="Enter extra info if you want">
          </div>
          <div class="form-group">
            <label for="length">max length</label>
            <select type="number" class="form-control" id="maxLength_In_PassphraseSetting" aria-describedby="max length of new passphrase string">
				<!--
				<option>128</option>
				<option>64</option>
				<option>56</option>
				<option>40</option>
				<option>32</option>
				<option>24</option>
				<option>16</option>
				-->
			</select>
          </div>
          <div class="form-group">
            <label><input type="checkbox" name="UseSymbol_In_PassphraseSetting" value="true" id="UseSymbol_In_PassphraseSetting">Use symbol in passphrase</label>
          </div>
          <button type="submit" class="btn btn-primary" id="SubmitSaveSetting"  >Save setting</button>
        </div><!--class="card-body-->
      </div><!--<div class="card"> card for  passphrase information-->
    </div><!-- [end] container for passphrase information on left side -->
  </div>
  <div class="col-md-4">
    <div class="container"><!-- [start] container for user information right side -->
      <div class="card"><!--card for table information-->
        <h4 class="card-title"> Passphrase Settings</h4>
        <div class="card-body"><!--<div class="card"> card for table-->
          <table class="table table-striped">
            <thead>
              <th scope="col">Passphrase Settings</th>
            </thead>
            <button type="submit" class="btn btn-primary" id="SubmitGetSetting_In_PassphraseSettings" >Get setting</button>
            <tbody id="user_passphrase_settings_In_PassphraseSettings">
              <!-- ここも設定一覧で収める。ただしユーザー毎に設定値が異なるので、jsで取得させる-->
                <tr><td id=XXX_In_PassphraseSetting algorithm=XXX, extra=XXX>Column content</td></tr>
                <tr><td id=XXX_In_PassphraseSetting algorithm=XXX, extra=XXX>Column content</td></tr>
                <tr><td id=XXX_In_PassphraseSetting algorithm=XXX, extra=XXX>Column content</td></tr>
                <tr><td id=XXX_In_PassphraseSetting algorithm=XXX, extra=XXX>Column content</td></tr>
                <tr><td id=XXX_In_PassphraseSetting algorithm=XXX, extra=XXX>Column content</td></tr>
            </tbody>
          </table>
        </div><!--<div class="card"> card for table-->
      </div><!--<div class="card"> card for table information-->
    </div><!-- [end] container for passphrase information on left side -->
  </div>
  <div class="col-md-4">
    <div class="container"><!-- [start] container for user information right side -->
      <div class="card"><!--card for creating login user-->
        <h4 class="card-title"> User account</h4>
        <div class="card-body">
          <div class="form-group">
            <div class="radio-area">
              <div class="form-group">
                <input type="radio" name="SelectOperation_In_UserAccount" value="Create User" checked="checked" onChange="createUserForm()">
                <label>Create User</label>
                <script>
                  function createUserForm() {
                    document.getElementById("UsernameForm_In_UserAccount").style.display="block";
                    document.getElementById("LoginPassphraseForm_In_UserAccount").style.display="block";
                    document.getElementById("SubmitUser_In_UserAccount").innerText="Create";
                  }
                </script>
                <input type="radio" name="SelectOperation_In_UserAccount" value="Update User" onChange="updateUserForm()">
                <label>Update User</label>
                <script>
                  function updateUserForm() {
                    document.getElementById("UsernameForm_In_UserAccount").style.display="none";
                    document.getElementById("LoginPassphraseForm_In_UserAccount").style.display="block";
                    document.getElementById("SubmitUser_In_UserAccount").innerText="Update";
                  }
                </script>
                <input type="radio" name="SelectOperation_In_UserAccount" value="Create User" onChange="deleteUserForm()">
                <label>Delete User</label>
                <script>
                  function deleteUserForm() {
                    document.getElementById("UsernameForm_In_UserAccount").style.display="none";
                    document.getElementById("LoginPassphraseForm_In_UserAccount").style.display="none";
                    document.getElementById("SubmitUser_In_UserAccount").innerText="Delete";
                  }
                </script>
              </div>
            </div>
            <!--
            <select class="col-xs-3" id="SelectOperation_In_UserAccount" name="SelectOperation_In_UserAccount" onChange="updateForm()">
              <option>Create</option>
              <option>Update</option>
              <option>Delete</option>
            </select>-->
          </div>
          <div class="form-group" id="UsernameForm_In_UserAccount">
            <label for="username">Username</label>
            <input type="text" class="form-control" id="Username_In_UserAccount" placeholder="Enter username">
          </div>
          <div class="form-group" id="LoginPassphraseForm_In_UserAccount">
            <label for="passphrase">Login passphrase</label>
            <input type="password" class="form-control" id="LoginPassphrase_In_UserAccount" placeholder="Enter passphrase">
          </div>
          <button type="submit" class="btn btn-primary" id="SubmitUser_In_UserAccount" >Create</button>
        </div>
        <div class="form-group">
          <button type="submit" class="btn btn-primary" id="SubmitLogoutUser_In_UpdateUser" >Logout</button>
        </div>
      <div><!--<div class="card"> card for login information-->
    </div><!-- [end] container for user information right side -->
  </div>
</div>
</body>
