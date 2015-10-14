package controller

// Perform Setup Here
func TestMain(m *testing.M) {
	HydrateTestData(&database)
	os.Exit(m.Run())
}

//----------
// POST Test
//----------

func TestValidPostingData {

}

func TestPostingNoData {

}

func TestPostingInvalidData {

}

//----------
// GET Test
//----------

func TestGettingValidObject {

}

func TestGettingUnknownObject {

}

func TestGettingObjectWithInvalidID {

}

//------------
// UPDATE Test
//------------

func TestUpdatingObject {

}

func TestUpdatingUnknownObject {

}

func TestUpdatingObjectWithBadData {

}

//------------
// DELET Test
//------------

func TestDeletingObject {

}

fun TestDeletingUnknownObject {

}