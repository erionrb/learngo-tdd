package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Should find known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Should not find unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Should add new word", func(t *testing.T) {
		dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertNull(t, err, "Should find added word")
		assertStrings(t, got, want)
	})

	t.Run("Should not repeat existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {

	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	t.Run("Should update to new definition", func(t *testing.T) {
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		got, err := dictionary.Search(word)

		assertNull(t, err, "Should find updated word")
		assertStrings(t, got, newDefinition)
	})

	t.Run("Should update when definition is same", func(t *testing.T) {
		newDefinition := "this is just a test"

		dictionary.Update(word, newDefinition)
		got, err := dictionary.Search(word)
		assertNull(t, err, "Should find updated word")
		assertStrings(t, got, newDefinition)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNull(t testing.TB, got interface{}, message string) {
	t.Helper()

	if got != nil {
		t.Fatal(message)
	}
}
