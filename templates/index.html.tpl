<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <script type="text/javascript" src="js/main.js"></script>
    <title>Generate passphrase</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <!--for css-->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
    <!--for js-->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js" integrity="sha384-wHAiFfRlMFy6i5SRaxvfOCifBUQy1xHdJ/yoi7FRNXMRBu5WHdZYu1hA6ZOblgut" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js" integrity="sha384-B0UglyR+jN6CkvvICOB2joaf5I4l3gm9GU6Hc1og6Ls7i6U/mkkaduKaBhlAXv9k" crossorigin="anonymous"></script>
  </head>
<body>
<div class="container">
<nav class="navbar navbar-light bg-light">
  <a class="navbar-brand display-1" href="https://github.com/developer-kikikaikai/CreatingHashPassword" data-toggle="tooltip" data-placement="left" data-original-title="Go Github repository to click here!">
    <img src="images/logo.png" height="60" class="d-inline-block align-top" alt="">Generate passphrase
  </a>
      <h4>-いつものパスワードから複雑なpassphraseを生成しよう！-</h4>
  <button align="right" type="submit" class="btn btn-primary pull-right" id="SubmitLogoutUser_In_UserAccount" >Logout</button>
</nav>
<ul class="nav nav-tabs">
  <li class="nav-item active"><a href="#passphrase_tab" class="nav-link" data-toggle="tab">Passphrase</a></li>
  <li class="nav-item"><a href="#account_tab" class="nav-link" data-toggle="tab">Account</a></li>
</ul>
<div class="tab-content">
  <div class="tab-pane well active" id="passphrase_tab">
    <div class="p-3">
      <div class="row">
        <div class="col-md-6">
          <div class="card"><!-- card for Generate Passphrase-->

          <div class="card border-success"><!-- card for Result-->
            <div class="card-body">
              <div class="form-group">
                <h1 class="card-title">Passphrase</h1>
                <input type="text" class="form-control form-control-lg" id="result_In_Generated_passphrase" value="" readonly>
              </div>
              <div class="form-group">
                <h5 class="card-title">How to generate passphrase</h5>
                <li> Enter setting in <strong>Seed for Passphrase</strong></li>
                <li> Enter Keyphrase and push button in <strong>Generate Passphrase</strong></li>
              </div>
            </div>
          </div><!-- card for Generate Passphrase-->

          <div class="card">
            <h2 class="card-header">Generate Passphrase</h2>
            <div class="card"><!-- card for Generate in Generate Passphrase-->
              <div class="card-body">
                <h4 class="card-title">Generate</h4>
                <div class="form-group">
                  <label for="inputKeyphrase">Keyphrase</label>
                  <input type="password" class="form-control" id="keyphrase_In_PassphraseSetting" placeholder="Keyphrase">
                </div>
                <button type="submit" class="btn btn-primary" id="SubmitGeneratePassphrase" data-toggle="tooltip" data-placement="left" data-original-title="Use 'Keyphrase' and 'Seed'!">Let's Generate!</button>
              </div>
            </div><!--<div class="card"> card for Generate in Generate Passphrase-->
          </div>
          </div>
        </div><!-- <div class="col-md-6"> -->
        <div class="col-sm-6" id="GroupPassphraseSettings">
          <div class="card"><!--card for Seed for Passphrase-->
            <h1 class="card-header">Seed for Passphrase</h1>
            <div class="card"><!--card for Current Setting in Seed for Passphrase-->
              <div class="card-body">
                <h4 class="card-title">Current Setting</h4>
                <div class="form-group">
                  <label for="inputTitle">Title</label>
                  <input type="text" class="form-control" id="title_In_PassphraseSetting" aria-describedby="title" value="Enter title">
                </div>
                <div class="form-group">
                  <label for="algorithm">Algorithm</label>
                  <select class="form-control" id="algorithmSelect_In_PassphraseSetting">
                    {{range $index, $algorithm := .Algorithms}}
                    <option>{{$algorithm}}</option>
                    {{end}}
                  </select>
                </div>
                <div class="form-group">
                  <label for="extraInfo">Extra</label>
                  <input type="text" class="form-control" id="extraInfo_In_PassphraseSetting" aria-describedby="extraInfo to get new passphrase string" placeholder="Enter extra info if you want">
                </div>
                <div class="form-group">
                  <label for="length">Length</label>
                  <select type="number" class="form-control" id="maxLength_In_PassphraseSetting" aria-describedby="max length of new passphrase string">
          	      </select>
                </div>
                <div class="form-group">
                  <label><input type="checkbox" name="UseSymbol_In_PassphraseSetting" value="true" id="UseSymbol_In_PassphraseSetting">Contain Symbol</label>
                </div>
                <button type="submit" class="btn btn-primary" id="SubmitSaveSetting" data-toggle="tooltip" data-placement="left" data-original-title="Save this setting to server!" >Save It!</button>
              </div><!--class="card-body-->
            </div><!--<div class="card"> card for Current Setting in Seed for Passphrase-->
            <div class="card"><!--card for table information-->
              <div class="card-body"><!--<div class="card"> card for table-->
                <h4 class="card-title"> Your Settings </h4>
                <table class="table table-striped">
                  <tbody id="user_passphrase_settings_In_PassphraseSettings">
                  </tbody>
                </table>
                <button type="submit" class="btn btn-primary" id="SubmitGetSetting_In_PassphraseSettings" data-toggle="tooltip" data-placement="left" data-original-title="Get your settings from server!">Get Settings!</button>
              </div><!--<div class="card"> card for table-->
            </div><!--<div class="card"> card for table information-->
          </div><!--<div class="card">card for Seed for Passphrase-->
        </div><!-- <div class="col-md-6" style="display: none" id="GroupPassphraseSettings"> -->
      </div><!-- <div class="row"> -->
    </div><!-- <div class="p-3"> -->
  </div> <!-- <div class="tab-pane active" id="passphrase_tab"> -->
  <div class="tab-pane" id="account_tab">
    <div class="p-3">
      <div class="col-md-6">
        <div class="card"><!--card for User Account-->
          <h1 class="card-header">User Account</h1>
          <div class="card"><!-- card for Operation in User Account-->
            <div class="card-body">
              <h4 class="card-title"> Operation </h4>
              <div class="form-group">
                <div class="radio-area">
                  <input type="radio" name="SelectOperation_In_UserAccount" value="Create User" checked="checked" onChange="createUserForm()">
                  <label>Create</label>
                  <script>
                    function createUserForm() {
                      document.getElementById("Username_In_UserAccount").disabled=false;
                      document.getElementById("LoginPassphrase_In_UserAccount").disabled=false;
                      document.getElementById("SubmitUser_In_UserAccount").innerText="Create Account!";
                    }
                  </script>
                  <input type="radio" name="SelectOperation_In_UserAccount" value="Update User" onChange="updateUserForm()">
                  <label>Update</label>
                  <script>
                    function updateUserForm() {
                      document.getElementById("Username_In_UserAccount").disabled=true;
                      document.getElementById("LoginPassphrase_In_UserAccount").disabled=false;
                      document.getElementById("SubmitUser_In_UserAccount").innerText="Update Passphrase!";
                    }
                  </script>
                  <input type="radio" name="SelectOperation_In_UserAccount" value="Create User" onChange="deleteUserForm()">
                  <label>Delete</label>
                  <script>
                    function deleteUserForm() {
                      document.getElementById("Username_In_UserAccount").disabled=true;
                      document.getElementById("LoginPassphrase_In_UserAccount").disabled=true;
                      document.getElementById("SubmitUser_In_UserAccount").innerText="Delete Account!";
                    }
                  </script>
                </div><!-- <div class="radio-area"> -->
              </div>
            </div>
          </div><!-- <div class="card">card for Operation in User Account-->
          <div class="card"><!-- card for Account-->
            <div class="card-body">
              <h4 class="card-title"> Account </h4>
              <div class="form-group" id="UsernameForm_In_UserAccount">
                <label for="username">Username</label>
                <input type="text" class="form-control" id="Username_In_UserAccount" placeholder="Enter username">
              </div>
              <div class="form-group" id="LoginPassphraseForm_In_UserAccount">
                <label for="passphrase">Passphrase</label>
                <input type="password" class="form-control" id="LoginPassphrase_In_UserAccount" placeholder="Enter passphrase">
              </div>
              <button type="submit" class="btn btn-primary" id="SubmitUser_In_UserAccount" >Create account!</button>
            </div>
          </div>
        </div>
      </div><!-- <div class="p-3"> -->
    </div><!-- <div class="tab-pane" id="account_tab"> -->
  </div><!---<div class="tab-content">-->
</div>
  <script>
    $('[data-toggle="tooltip"]').tooltip();
  </script>
</body>
