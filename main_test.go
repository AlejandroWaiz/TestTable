package main

import "testing"

func TestDivide(t *testing.T) {

	//Declaramos un array de structs. De esta forma podemos elegir de antemano los valores que usaremos en el test y asi
	//poder crear distintos casos sin tener que reescribir variables a medida que las funciones se hagan mas grandes.
	testTable := []struct {
		testName                           string
		firstValue, secondValue, expResult int
		wantErr                            bool
		expErr                             error
	}{
		//Creamos variados tests para cada posibilidad que queramos-.
		{testName: "Normal Divide", firstValue: 10, secondValue: 5, expResult: 2},
		{testName: "Divide by zero", firstValue: 10, secondValue: 0, expResult: 0, wantErr: true, expErr: &ErrDivideByZero{}},
	}

	//Aca iteramos en la tabla de tests, utilizando un valor unico "testCase"
	for _, testCase := range testTable {

		//Usamos t.Run("NombreDeTuCasoDeTest", func(t *testing.T ){ <Aca va tu funcion de test>}) para poder correr
		//''subtests'' durante la iteracion, cada uno designado por su nombre en la tabla. Pruebalo! go test -v
		t.Run(testCase.testName, func(t *testing.T) {

			obtainedResult, obtainedErr := Divide(testCase.firstValue, testCase.secondValue)

			//Aca vemos si estamos esperando un error. Si hubo uno y no lo estabamos esperando, hacemos fallar el test.
			if testCase.wantErr == false && obtainedErr != nil {
				t.Errorf("Wasnt expecting error, got: %v", obtainedErr)
			}

			//Si estabamos esperando un error y es distinto al que obtuvimos, hacemos fallar el test.
			if testCase.expErr != obtainedErr {
				t.Errorf("Expecting: %v, got: %v", testCase.expErr, obtainedErr)
			}

			//Si el resultado que obtuvimos es distinto a lo que esperamos, hacemos fallar el test.
			if obtainedResult != testCase.expResult {
				t.Errorf("Expecting %v as  result, got %v", testCase.expResult, obtainedResult)
			}

			//Aca termina el flujo de la iteración para pasar a la siguiente.

		})

	}

}
