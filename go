<!DOCTYPE html>
<html>
<head>
</head>
<body>
	<form method ="post">
		ID No. <input type="number" name="id"><br>
		Name. <input type="text" name=name><br>
		Age. <input type="number" name="age"><br>
		Gender. <input type="text" name="gender"><br>
		<input type="submit" value="save" name="btnSave"> 
		<input type="submit" value="update" name="btnUpdate">
		<input type="submit" value="update" name="btnDelete">
	</form>
	<?php
	
	$sql="SELECT * FROM tblInfo";
	echo "<table border='1'>";
	echo "<tr>";
	echo "<th>ID No</th>";
	echo "<th>Name</th>";
	echo "<th>Age</th>";
	echo "<th>Gender</th>";
	echo "</tr>";
	if(isset($result)>0)
	{
	while($row= mysqli_fetch_array($result))
		$id=$row['id'];
		$name=$row['name'];
		$age=$row['age'];
		$gender=$row['gender'];
		echo "<tr>";
		echo "<td> $id </td>
			  <td> $name </td>
			  <td> $age </td>
			  <td> $gender</td>";
		echo "</tr>";
	}
echo "</table>";
?>
</body>
</html>
