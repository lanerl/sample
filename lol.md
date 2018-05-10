CONTROLLER

<?php
class Reviews extends CI_Controller {
}

public function show($id) {
$this->load->model('reviews_model');
$tblInfo = $this->reviews_model->get_tblInfo($id);

public $id;
public $name;
public $age;
public $gender;

public function insert()
	$this->id = $_POST['id'];
	$this->name = $_POST['name'];
	$this->age = $_POST['age'];
	$this->gender = $_POST['gender'];

$this->load->view('movie_review', $data);
}

?>

============================================\
MODEL

<?php
class News_model extends CI_Model {
}
public function __construct() {
$this->load->database('test');
}

public function get_reviews($id) {
if($id != FALSE) {
$query = $this->db->get_where('tblInfo', array('id' => $id));
return $query->row_array();
}
else {
return FALSE;
}
}

		$server = "localhost";
		$username="root";
		$password="";
		$dbname="test";
		$conn = new mysqli ($server,$username,$password,$dbname);
		if($conn->connect_error)
		{
			echo "Cannot connect to database server!";
		} 

		if(isset($_POST['btnSave']))
{
		$id = $_POST['id'];
		$name = $_POST['name'];
		$age = $_POST['age'];
		$gender = $_POST['gender'];
		$sql ="INSERT INTO tblInfo VALUES($id, '$name', $age, '$gender')";
		$result = $conn->query($sql);

	if($result == TRUE){
		echo "<script>alert('Added!');location.replace('movie_review.php')</script>";
}	else{
		echo "Problem in saving record!";
}
}

if(isset($_POST['btnUpdate']))
{
	$id = $_POST['id'];
	$name = $_POST['name'];
	$age = $_POST['age'];
	$gender = $_POST['gender'];
	$sql = "UPDATE tblInfo SET name='$name' , age=$age ,gender='$gender' WHERE id=$id";
	$result = $conn->query($sql);

	if($result==TRUE){
		echo "<script>alert('Updated!');location.replace('info2.php')</script>";
	}	else{

			echo "Problem in updating record!";
	}
}


if(isset($_POST['btnDelete']))
{
	$id = $_POST['id'];
	$name = $_POST['name'];
	$age = $_POST['age'];
	$gender = $_POST['gender'];
	$sql = "DELETE FROM tblInfo SET name='$name' , age=$age ,gender='$gender' WHERE id=$id";
	$result = $conn->query($sql);

	if($result==TRUE){
		echo "<script>alert('Updated!');location.replace('info2.php')</script>";
	}	else{

			echo "Problem in updating record!";
	}
}


?>
==========================================
VIEWS

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
	$result = $conn->query($sql);
	echo "<table border='1'>";
	echo "<tr>";
	echo "<th>ID No</th>";
	echo "<th>Name</th>";
	echo "<th>Age</th>";
	echo "<th>Gender</th>";
	echo "</tr>";
	while($row=$result->fetch_array())
	{
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
