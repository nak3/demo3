<!DOCTYPE html>
<html>
<head>
<meta charset='utf-8'>
<title>OpenShift v3 demonstration</title>
</head>
<body>
<h1>
<?php
  $kakushi = htmlspecialchars($_GET['kakushi']);
  echo "{$kakushi}";
?>
</h1>
<form action="index.php" method="get">
  <input type="hidden" name="kakushi" value="Morning OpenShift v3!">
  <input type="submit">
</form>

</body>
</html>
