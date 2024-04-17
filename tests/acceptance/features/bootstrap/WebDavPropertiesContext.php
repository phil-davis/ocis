<?php declare(strict_types=1);
/**
 * ownCloud
 *
 * @author Artur Neumann <artur@jankaritech.com>
 * @copyright Copyright (c) 2019, ownCloud GmbH
 *
 * This code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License,
 * as published by the Free Software Foundation;
 * either version 3 of the License, or any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>
 *
 */

use Behat\Behat\Context\Context;
use Behat\Behat\Hook\Scope\BeforeScenarioScope;
use Behat\Gherkin\Node\TableNode;
use PHPUnit\Framework\Assert;
use TestHelpers\Asserts\WebDav as WebDavTest;
use TestHelpers\WebDavHelper;
use Psr\Http\Message\ResponseInterface;

require_once 'bootstrap.php';

/**
 * Steps that relate to managing file/folder properties via WebDav
 */
class WebDavPropertiesContext implements Context {
	private FeatureContext $featureContext;

	/**
	 * @var array map with user as key and another map as value,
	 *            which has path as key and etag as value
	 */
	private array $storedETAG = [];

	/**
	 * @When /^user "([^"]*)" gets the properties of (?:file|folder|entry) "([^"]*)" using the WebDAV API$/
	 *
	 * @param string $user
	 * @param string $path
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userGetsThePropertiesOfFolder(
		string $user,
		string $path
	):void {
		$response = $this->featureContext->listFolder(
			$user,
			$path,
			'0'
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @When /^user "([^"]*)" gets the properties of (?:file|folder|entry) "([^"]*)" with depth (\d+) using the WebDAV API$/
	 *
	 * @param string $user
	 * @param string $path
	 * @param string $depth
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userGetsThePropertiesOfFolderWithDepth(
		string $user,
		string $path,
		string $depth
	):void {
		$response = $this->featureContext->listFolder(
			$user,
			$path,
			$depth
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @param string $user
	 * @param string $path
	 * @param TableNode|null $propertiesTable
	 *
	 * @return ResponseInterface
	 * @throws Exception
	 */
	public function getPropertiesOfFolder(
		string $user,
		string $path,
		TableNode $propertiesTable
	): ResponseInterface {
		$user = $this->featureContext->getActualUsername($user);
		$properties = null;
		$this->featureContext->verifyTableNodeColumns($propertiesTable, ["propertyName"]);
		$this->featureContext->verifyTableNodeColumnsCount($propertiesTable, 1);
		foreach ($propertiesTable->getColumnsHash() as $row) {
			$properties[] = $row["propertyName"];
		}
		return $this->featureContext->listFolder(
			$user,
			$path,
			"1",
			$properties
		);
	}

	/**
	 * @When /^user "([^"]*)" gets the following properties of (?:file|folder|entry) "([^"]*)" using the WebDAV API$/
	 *
	 * @param string $user
	 * @param string $path
	 * @param TableNode|null $propertiesTable
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userGetsFollowingPropertiesOfEntryUsingWebDavApi(
		string $user,
		string $path,
		TableNode $propertiesTable
	):void {
		$response = $this->getPropertiesOfFolder($user, $path, $propertiesTable);
		$this->featureContext->setResponse($response);
		$this->featureContext->pushToLastStatusCodesArrays();
	}

	/**
	 * @When /^the user gets the following properties of (?:file|folder|entry) "([^"]*)" using the WebDAV API$/
	 *
	 * @param string $path
	 * @param TableNode|null $propertiesTable
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theUserGetsPropertiesOfFolder(string $path, TableNode $propertiesTable) {
		$response = $this->getPropertiesOfFolder(
			$this->featureContext->getCurrentUser(),
			$path,
			$propertiesTable
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @Given /^user "([^"]*)" has set the following properties of (?:file|folder|entry) "([^"]*)" using the WebDav API$/
	 *
	 * if no namespace prefix is provided before property, default `oc:` prefix is set for added props
	 * only and everything rest on xml is set to prefix `d:`
	 *
	 * @param string $username
	 * @param string $path
	 * @param TableNode|null $propertiesTable with following columns with column header as:
	 *                                        property: name of prop to be set
	 *                                        value: value of prop to be set
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userHasSetFollowingPropertiesUsingProppatch(string $username, string $path, TableNode $propertiesTable) {
		$username = $this->featureContext->getActualUsername($username);
		$this->featureContext->verifyTableNodeColumns($propertiesTable, ['propertyName', 'propertyValue']);
		$properties = $propertiesTable->getColumnsHash();
		$response = WebDavHelper::proppatchWithMultipleProps(
			$this->featureContext->getBaseUrl(),
			$username,
			$this->featureContext->getPasswordForUser($username),
			$path,
			$properties,
			$this->featureContext->getStepLineRef(),
			$this->featureContext->getDavPathVersion()
		);
		$this->featureContext->theHTTPStatusCodeShouldBe(207, "", $response);
	}

	/**
	 * @When user :user gets a custom property :propertyName with namespace :namespace of file :path
	 *
	 * @param string $user
	 * @param string $propertyName
	 * @param string $namespace namespace in form of "x1='http://whatever.org/ns'"
	 * @param string $path
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userGetsPropertiesOfFile(
		string $user,
		string $propertyName,
		string $namespace,
		string $path
	):void {
		$user = $this->featureContext->getActualUsername($user);
		$properties = [
			$namespace => $propertyName
		];
		$this->featureContext->setResponse(
			WebDavHelper::propfind(
				$this->featureContext->getBaseUrl(),
				$this->featureContext->getActualUsername($user),
				$this->featureContext->getUserPassword($user),
				$path,
				$properties,
				$this->featureContext->getStepLineRef(),
				"0",
				"files",
				$this->featureContext->getDavPathVersion()
			)
		);
	}

	/**
	 * @param string $path
	 * @param TableNode $propertiesTable
	 *
	 * @return ResponseInterface
	 * @throws Exception
	 */
	public function getPropertiesOfEntryFromLastLinkShare(string $path, TableNode $propertiesTable): ResponseInterface {
		$user = $this->featureContext->getLastCreatedPublicShareToken();
		$properties = null;
		foreach ($propertiesTable->getRows() as $row) {
			$properties[] = $row[0];
		}
		return $this->featureContext->listFolder(
			$user,
			$path,
			'0',
			$properties,
			$this->featureContext->getDavPathVersion() === 1 ? "public-files" : "public-files-new"
		);
	}

	/**
	 * @When /^the public gets the following properties of (?:file|folder|entry) "([^"]*)" in the last created public link using the WebDAV API$/
	 *
	 * @param string $path
	 * @param TableNode $propertiesTable
	 *
	 * @return void
	 * @throws Exception
	 */
	public function thePublicGetsFollowingPropertiesOfEntryFromLastLinkShare(string $path, TableNode $propertiesTable):void {
		$response = $this->getPropertiesOfEntryFromLastLinkShare($path, $propertiesTable);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @param string $user user id who sets the property
	 * @param string $propertyName name of property in Clark notation
	 * @param string $namespace namespace in form of "x1='http://whatever.org/ns'"
	 * @param string $path path on which to set properties to
	 * @param string $propertyValue property value
	 *
	 * @return ResponseInterface
	 * @throws Exception
	 */
	public function setPropertyWithNamespaceOfResource(
		string $user,
		string $propertyName,
		string $namespace,
		string $path,
		string $propertyValue
	):ResponseInterface {
		$user = $this->featureContext->getActualUsername($user);
		return WebDavHelper::proppatch(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getActualUsername($user),
			$this->featureContext->getUserPassword($user),
			$path,
			$propertyName,
			$propertyValue,
			$this->featureContext->getStepLineRef(),
			$namespace,
			$this->featureContext->getDavPathVersion()
		);
	}

	/**
	 * @When /^user "([^"]*)" sets property "([^"]*)"  with namespace "([^"]*)" of (?:file|folder|entry) "([^"]*)" to "([^"]*)" using the WebDAV API$/
	 *
	 * @param string $user user id who sets the property
	 * @param string $propertyName name of property in Clark notation
	 * @param string $namespace namespace in form of "x1='http://whatever.org/ns'"
	 * @param string $path path on which to set properties to
	 * @param string $propertyValue property value
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userSetsPropertyWithNamespaceOfEntryTo(
		string $user,
		string $propertyName,
		string $namespace,
		string $path,
		string $propertyValue
	):void {
		$response = $this->setPropertyWithNamespaceOfResource(
			$user,
			$propertyName,
			$namespace,
			$path,
			$propertyValue
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @Given /^user "([^"]*)" has set property "([^"]*)" with namespace "([^"]*)" of (?:file|folder|entry) "([^"]*)" to "([^"]*)"$/
	 *
	 * @param string $user user id who sets the property
	 * @param string $propertyName name of property in Clark notation
	 * @param string $namespace namespace in form of "x1='http://whatever.org/ns'"
	 * @param string $path path on which to set properties to
	 * @param string $propertyValue property value
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userHasSetPropertyWithNamespaceOfEntryTo(
		string $user,
		string $propertyName,
		string $namespace,
		string $path,
		string $propertyValue
	):void {
		$response = $this->setPropertyWithNamespaceOfResource(
			$user,
			$propertyName,
			$namespace,
			$path,
			$propertyValue
		);
		$this->featureContext->theHTTPStatusCodeShouldBe(207, "", $response);
	}

	/**
	 * @Then /^the response should contain a custom "([^"]*)" property with namespace "([^"]*)" and value "([^"]*)"$/
	 *
	 * @param string $propertyName
	 * @param string $namespaceString
	 * @param string $propertyValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theResponseShouldContainACustomPropertyWithValue(
		string $propertyName,
		string $namespaceString,
		string $propertyValue
	):void {
		$responseXmlObject = $this->featureContext->getResponseXml(
			$this->featureContext->getResponse(),
			__METHOD__
		);
		//calculate the namespace prefix and namespace
		$matches = [];
		\preg_match("/^(.*)='(.*)'$/", $namespaceString, $matches);
		$nameSpace = $matches[2];
		$nameSpacePrefix = $matches[1];
		$responseXmlObject->registerXPathNamespace(
			$nameSpacePrefix,
			$nameSpace
		);
		$xmlPart = $responseXmlObject->xpath(
			"//d:prop/" . "$nameSpacePrefix:$propertyName"
		);
		Assert::assertArrayHasKey(
			0,
			$xmlPart,
			"Cannot find property \"$propertyName\""
		);
		Assert::assertEquals(
			$propertyValue,
			$xmlPart[0]->__toString(),
			"\"$propertyName\" has a value \"" .
			$xmlPart[0]->__toString() . "\" but \"$propertyValue\" expected"
		);
	}

	/**
	 * @Then /^the response should contain a custom "([^"]*)" property with namespace "([^"]*)" and complex value "(([^"\\]|\\.)*)"$/
	 *
	 * @param string $propertyName
	 * @param string $namespaceString
	 * @param string $propertyValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theResponseShouldContainACustomPropertyWithComplexValue(
		string $propertyName,
		string $namespaceString,
		string $propertyValue
	):void {
		// let's unescape quotes first
		$propertyValue = \str_replace('\"', '"', $propertyValue);
		$responseXmlObject = $this->featureContext->getResponseXml(
			$this->featureContext->getResponse(),
			__METHOD__
		);
		//calculate the namespace prefix and namespace
		$matches = [];
		\preg_match("/^(.*)='(.*)'$/", $namespaceString, $matches);
		$nameSpace = $matches[2];
		$nameSpacePrefix = $matches[1];
		$responseXmlObject->registerXPathNamespace(
			$nameSpacePrefix,
			$nameSpace
		);
		$xmlPart = $responseXmlObject->xpath(
			"//d:prop/" . "$nameSpacePrefix:$propertyName" . "/*"
		);
		Assert::assertArrayHasKey(
			0,
			$xmlPart,
			"Cannot find property \"$propertyName\""
		);
		Assert::assertEquals(
			$propertyValue,
			$xmlPart[0]->asXML(),
			"\"$propertyName\" has a value \"" .
			$xmlPart[0]->asXML() . "\" but \"$propertyValue\" expected"
		);
	}

	/**
	 * @Then the single response should contain a property :property with a child property :childProperty
	 *
	 * @param string $property
	 * @param string $childProperty
	 *
	 * @return void
	 *
	 * @throws Exception
	 */
	public function theSingleResponseShouldContainAPropertyWithChildProperty(
		string $property,
		string $childProperty
	):void {
		$xmlPart = $this->featureContext->getResponseXml()->xpath(
			"//d:prop/$property/$childProperty"
		);
		Assert::assertTrue(
			isset($xmlPart[0]),
			"Cannot find property \"$property/$childProperty\""
		);
	}

	/**
	 * @Then the single response should contain a property :key with value :value
	 *
	 * @param string $key
	 * @param string $expectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theSingleResponseShouldContainAPropertyWithValue(
		string $key,
		string $expectedValue
	):void {
		$this->checkResponseContainsAPropertyWithValue(
			$this->featureContext->getResponse(),
			$key,
			$expectedValue,
			$expectedValue
		);
	}

	/**
	 * @Then the single response about the file owned by :user should contain a property :key with value :value
	 *
	 * @param string $user
	 * @param string $key
	 * @param string $expectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theSingleResponseAboutTheFileOwnedByShouldContainAPropertyWithValue(
		string $user,
		string $key,
		string $expectedValue
	):void {
		$this->checkResponseContainsAPropertyWithValue(
			$this->featureContext->getResponse(),
			$key,
			$expectedValue,
			$expectedValue,
			$user
		);
	}

	/**
	 * @Then the single response should contain a property :key with value :value or with value :altValue
	 *
	 * @param string $key
	 * @param string $expectedValue
	 * @param string $altExpectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theSingleResponseShouldContainAPropertyWithValueAndAlternative(
		string $key,
		string $expectedValue,
		string $altExpectedValue
	):void {
		$this->checkResponseContainsAPropertyWithValue(
			$this->featureContext->getResponse(),
			$key,
			$expectedValue,
			$altExpectedValue
		);
	}

	/**
	 * @param ResponseInterface $response
	 * @param string $key
	 * @param string $expectedValue
	 * @param string $altExpectedValue
	 * @param string|null $user
	 *
	 * @return void
	 * @throws Exception
	 */
	public function checkResponseContainsAPropertyWithValue(
		ResponseInterface $response,
		string $key,
		string $expectedValue,
		string $altExpectedValue,
		?string $user = null
	):void {
		$xmlPart = $this->featureContext->getResponseXml($response);
		$xmlPart = $xmlPart->xpath(
			"//d:prop/$key"
		);
		Assert::assertTrue(
			isset($xmlPart[0]),
			"Cannot find property \"$key\""
		);
		$value = $xmlPart[0]->__toString();
		$expectedValue = $this->featureContext->substituteInLineCodes(
			$expectedValue,
			$user
		);
		$expectedValue = "#^$expectedValue$#";
		$altExpectedValue = "#^$altExpectedValue$#";
		if (\preg_match($expectedValue, $value) !== 1
			&& \preg_match($altExpectedValue, $value) !== 1
		) {
			Assert::fail(
				"Property \"$key\" found with value \"$value\", " .
				"expected \"$expectedValue\" or \"$altExpectedValue\""
			);
		}
	}

	/**
	 * @Then the value of the item :xpath in the response should be :value
	 *
	 * @param string $xpath
	 * @param string $expectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertValueOfItemInResponseIs(string $xpath, string $expectedValue):void {
		$this->assertValueOfItemInResponseAboutUserIs(
			$xpath,
			null,
			$expectedValue
		);
	}

	/**
	 * @Then as user :user the value of the item :xpath of path :path in the response should be :value
	 *
	 * @param string $user
	 * @param string $xpath
	 * @param string $path
	 * @param string $expectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function valueOfItemOfPathShouldBe(string $user, string $xpath, string $path, string $expectedValue):void {
		$path = $this->featureContext->substituteInLineCodes($path, $user);
		$fullXpath = "//d:response/d:href[.='$path']/following-sibling::d:propstat$xpath";
		$this->assertValueOfItemInResponseAboutUserIs(
			$fullXpath,
			null,
			$expectedValue
		);
	}

	/**
	 * @Then the value of the item :xpath in the response about user :user should be :value
	 *
	 * @param string $xpath
	 * @param string $user
	 * @param string $expectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theValueOfTheItemInTheResponseAboutUserShouldBe(string $xpath, string $user, string $expectedValue):void {
		$this->assertValueOfItemInResponseAboutUserIs($xpath, $user, $expectedValue);
	}

	/**
	 * @param string $xpath
	 * @param string|null $user
	 * @param string $expectedValue
	 *
	 * @return void
	 */
	public function assertValueOfItemInResponseAboutUserIs(string $xpath, ?string $user, string $expectedValue):void {
		$resXml = $this->featureContext->getResponseXml(
			$this->featureContext->getResponse(),
			__METHOD__
		);
		$value = $this->getXmlItemByXpath($resXml, $xpath);
		$user = $this->featureContext->getActualUsername($user);
		$expectedValue = $this->featureContext->substituteInLineCodes(
			$expectedValue,
			$user
		);

		// The expected value can contain /%base_path%/ which can be empty some time
		// This will result in urls such as //remote.php, so replace that
		$expectedValue = preg_replace("/\/\/remote\.php/i", "/remote.php", $expectedValue);
		Assert::assertEquals(
			$expectedValue,
			$value,
			"item \"$xpath\" found with value \"$value\", " .
			"expected \"$expectedValue\""
		);
	}

	/**
	 * @Then the value of the item :xpath in the response about user :user should be :value1 or :value2
	 *
	 * @param string $xpath
	 * @param string|null $user
	 * @param string $expectedValue1
	 * @param string $expectedValue2
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertValueOfItemInResponseAboutUserIsEitherOr(string $xpath, ?string $user, string $expectedValue1, string $expectedValue2):void {
		if (!$expectedValue2) {
			$expectedValue2 = $expectedValue1;
		}
		$resXml = $this->featureContext->getResponseXml(
			$this->featureContext->getResponse(),
			__METHOD__
		);
		$value = $this->getXmlItemByXpath($resXml, $xpath);
		$user = $this->featureContext->getActualUsername($user);
		$expectedValue1 = $this->featureContext->substituteInLineCodes(
			$expectedValue1,
			$user
		);

		$expectedValue2 = $this->featureContext->substituteInLineCodes(
			$expectedValue2,
			$user
		);

		// The expected value can contain /%base_path%/ which can be empty some time
		// This will result in urls such as //remote.php, so replace that
		$expectedValue1 = preg_replace("/\/\/remote\.php/i", "/remote.php", $expectedValue1);
		$expectedValue2 = preg_replace("/\/\/remote\.php/i", "/remote.php", $expectedValue2);
		$expectedValues = [$expectedValue1, $expectedValue2];
		$isExpectedValueInMessage = \in_array($value, $expectedValues);
		Assert::assertTrue($isExpectedValueInMessage, "The actual value \"$value\" is not one of the expected values: \"$expectedValue1\" or \"$expectedValue2\"");
	}

	/**
	 * @param <type> $xmlResponse
	 * @param string $xpath
	 *
	 * @return string
	 * @throws Exception
	 */
	public function getXmlItemByXpath(
		SimpleXMLElement $xmlResponse,
		string $xpath
	): string {
		$xmlPart = $xmlResponse->xpath($xpath);
		Assert::assertTrue(
			isset($xmlPart[0]),
			"Cannot find item with xpath \"$xpath\""
		);
		return $xmlPart[0]->__toString();
	}

	/**
	 * @Then the value of the item :xpath in the response should match :value
	 *
	 * @param string $xpath
	 * @param string $pattern
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertValueOfItemInResponseRegExp(string $xpath, string $pattern):void {
		$this->assertXpathValueMatchesPattern(
			$this->featureContext->getResponseXml(),
			$xpath,
			$pattern
		);
	}

	/**
	 * @Then /^as a public the lock discovery property "([^"]*)" of the (?:file|folder|entry) "([^"]*)" should match "([^"]*)"$/
	 *
	 * @param string $xpath
	 * @param string $path
	 * @param string $pattern
	 *
	 * @return void
	 * @throws Exception
	 */
	public function publicGetsThePropertiesOfFolderAndAssertValueOfItemInResponseRegExp(string $xpath, string $path, string $pattern):void {
		$propertiesTable = new TableNode([['propertyName'],['d:lockdiscovery']]);
		$response = $this->getPropertiesOfEntryFromLastLinkShare($path, $propertiesTable);
		$this->featureContext->theHTTPStatusCodeShouldBe('207', "", $response);
		$this->assertXpathValueMatchesPattern(
			$this->featureContext->getResponseXml($response),
			$xpath,
			$pattern
		);
	}

	/**
	 * @Then there should be an entry with href containing :expectedHref in the response to user :user
	 *
	 * @param string $expectedHref
	 * @param string $user
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertEntryWithHrefMatchingRegExpInResponseToUser(string $expectedHref, string $user):void {
		$resXml = $this->featureContext->getResponseXml(
			$this->featureContext->getResponse(),
			__METHOD__
		);

		$user = $this->featureContext->getActualUsername($user);
		$expectedHref = $this->featureContext->substituteInLineCodes(
			$expectedHref,
			$user,
			['preg_quote' => ['/']]
		);

		$index = 0;
		while (true) {
			$index++;
			$xpath = "//d:response[$index]/d:href";
			$xmlPart = $resXml->xpath($xpath);
			// If we have run out of entries in the response, then fail the test
			Assert::assertTrue(
				isset($xmlPart[0]),
				"Cannot find any entry having href with value $expectedHref in response to $user"
			);
			$value = $xmlPart[0]->__toString();
			$decodedValue = \rawurldecode($value);
			// for folders, decoded value will be like: "/owncloud/core/remote.php/webdav/strängé folder/"
			// expected href should be like: "remote.php/webdav/strängé folder/"
			// for files, decoded value will be like: "/owncloud/core/remote.php/webdav/strängé folder/file.txt"
			// expected href should be like: "remote.php/webdav/strängé folder/file.txt"
			$explodeDecoded = \explode('/', $decodedValue);
			// get the first item of the expected href.
			// i.e. remote.php from "remote.php/webdav/strängé folder/file.txt"
			// or dav from "dav/spaces/%spaceid%/C++ file.cpp"
			$explodeExpected = \explode('/', $expectedHref);
			$remotePhpIndex = \array_search($explodeExpected[0], $explodeDecoded);
			if ($remotePhpIndex) {
				$explodedHrefPartArray = \array_slice($explodeDecoded, $remotePhpIndex);
				$actualHrefPart = \implode('/', $explodedHrefPartArray);
				if ($this->featureContext->getDavPathVersion() === WebDavHelper::DAV_VERSION_SPACES) {
					// for spaces webdav, space id is included in the href
					// space id from our helper is returned as d8c029e0\-2bc9\-4b9a\-8613\-c727e5417f05
					// so we've to remove "\" before every "-"
					$expectedHref = str_replace('\-', '-', $expectedHref);
					$expectedHref = str_replace('\$', '$', $expectedHref);
					$expectedHref = str_replace('\!', '!', $expectedHref);
				}
				if ($actualHrefPart === $expectedHref) {
					break;
				}
			}
		}
	}

	/**
	 * @param SimpleXMLElement $responseXml
	 * @param string $xpath
	 * @param string $pattern
	 * @param string|null $user
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertXpathValueMatchesPattern(SimpleXMLElement $responseXml, string $xpath, string $pattern, ?string $user = null): void {
		$xmlPart = $responseXml->xpath($xpath);
		Assert::assertTrue(
			isset($xmlPart[0]),
			"Cannot find item with xpath \"$xpath\""
		);
		$user = $this->featureContext->getActualUsername($user);
		$value = $xmlPart[0]->__toString();
		$pattern = $this->featureContext->substituteInLineCodes(
			$pattern,
			$user,
			['preg_quote' => ['/']],
			[
				[
					"code" => "%public_token%",
					"function" =>
					[$this->featureContext, "getLastCreatedPublicShareToken"],
					"parameter" => []
				],
			]
		);
		Assert::assertMatchesRegularExpression(
			$pattern,
			$value,
			"item \"$xpath\" found with value \"$value\", " .
			"expected to match regex pattern: \"$pattern\""
		);
	}

	/**
	 * @Then the value of the item :xpath in the response to user :user should match :value
	 *
	 * @param string $xpath
	 * @param string|null $user
	 * @param string $pattern
	 * @param SimpleXMLElement|null $responseXml
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theValueOfItemInResponseToUserShouldMatch(string $xpath, ?string $user, string $pattern, ?SimpleXMLElement $responseXml = null):void {
		$resXml = $this->featureContext->getResponseXml();
		$this->assertXpathValueMatchesPattern($resXml, $xpath, $pattern, $user);
	}

	/**
	 * @Then /^as user "([^"]*)" the lock discovery property "([^"]*)" of the (?:file|folder|entry) "([^"]*)" should match "([^"]*)"$/
	 *
	 * @param string|null $user
	 * @param string $xpath
	 * @param string $path
	 * @param string $pattern
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userGetsPropertiesOfFolderAndAssertValueOfItemInResponseToUserRegExp(string $user, string $xpath, string $path, string $pattern):void {
		$propertiesTable = new TableNode([['propertyName'],['d:lockdiscovery']]);
		$response = $this->getPropertiesOfFolder(
			$user,
			$path,
			$propertiesTable
		);
		$this->featureContext->theHTTPStatusCodeShouldBe('207', '', $response);
		$this->assertXpathValueMatchesPattern(
			$this->featureContext->getResponseXml($response),
			$xpath,
			$pattern,
			$user
		);
	}

	/**
	 * @Then the item :xpath in the response should not exist
	 *
	 * @param string $xpath
	 *
	 * @return void
	 * @throws Exception
	 */
	public function assertItemInResponseDoesNotExist(string $xpath):void {
		$xmlPart = $this->featureContext->getResponseXml()->xpath($xpath);
		Assert::assertFalse(
			isset($xmlPart[0]),
			"Found item with xpath \"$xpath\" but it should not exist"
		);
	}

	/**
	 * @Then /^as user "([^"]*)" (?:file|folder|entry) "([^"]*)" should contain a property "([^"]*)" with value "([^"]*)" or with value "([^"]*)"$/
	 * @Then /^as user "([^"]*)" (?:file|folder|entry) "([^"]*)" should contain a property "([^"]*)" with value "([^"]*)"$/
	 *
	 * @param string $user
	 * @param string $path
	 * @param string $property
	 * @param string $expectedValue
	 * @param string $altExpectedValue
	 *
	 * @return void
	 * @throws Exception
	 */
	public function asUserFolderShouldContainAPropertyWithValueOrWithValue(
		string $user,
		string $path,
		string $property,
		string $expectedValue,
		?string $altExpectedValue = null
	):void {
		$this->checkPropertyOfAFolder($user, $path, $property, $expectedValue, $altExpectedValue);
	}

	/**
	 * @param string $user
	 * @param string $path
	 * @param string $property
	 * @param string $expectedValue
	 * @param string|null $altExpectedValue
	 *
	 * @return void
	 */
	public function checkPropertyOfAFolder(
		string $user,
		string $path,
		string $property,
		string $expectedValue,
		?string $altExpectedValue = null
	):void {
		$response = $this->featureContext->listFolder(
			$user,
			$path,
			'0',
			[$property]
		);
		if ($altExpectedValue === null) {
			$altExpectedValue = $expectedValue;
		}
		$this->checkResponseContainsAPropertyWithValue(
			$response,
			$property,
			$expectedValue,
			$altExpectedValue
		);
	}

	/**
	 * @Then the single response should contain a property :key with value like :regex
	 *
	 * @param string $key
	 * @param string $regex
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theSingleResponseShouldContainAPropertyWithValueLike(
		string $key,
		string $regex
	):void {
		$xmlPart = $this->featureContext->getResponseXml()->xpath(
			"//d:prop/$key"
		);
		Assert::assertTrue(
			isset($xmlPart[0]),
			"Cannot find property \"$key\""
		);
		$value = $xmlPart[0]->__toString();
		Assert::assertMatchesRegularExpression(
			$regex,
			$value,
			"Property \"$key\" found with value \"$value\", expected \"$regex\""
		);
	}

	/**
	 * @Then the response should contain a share-types property with
	 *
	 * @param TableNode $table
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theResponseShouldContainAShareTypesPropertyWith(TableNode $table):void {
		$this->featureContext->verifyTableNodeColumnsCount($table, 1);
		WebdavTest::assertResponseContainsShareTypes(
			$this->featureContext->getResponseXml(),
			$table->getRows()
		);
	}

	/**
	 * @Then the response should contain an empty property :property
	 *
	 * @param string $property
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theResponseShouldContainAnEmptyProperty(string $property):void {
		$xmlPart = $this->featureContext->getResponseXml()->xpath(
			"//d:prop/$property"
		);
		Assert::assertCount(
			1,
			$xmlPart,
			"Cannot find property \"$property\""
		);
		Assert::assertEmpty(
			$xmlPart[0],
			"Property \"$property\" is not empty"
		);
	}

	/**
	 * @param string $user
	 * @param string $path
	 * @param string|null $storePath
	 *
	 * @return SimpleXMLElement
	 * @throws Exception
	 */
	public function storeEtagOfElement(string $user, string $path, ?string $storePath = ""):SimpleXMLElement {
		if ($storePath === "") {
			$storePath = $path;
		}
		$user = $this->featureContext->getActualUsername($user);
		$propertiesTable = new TableNode([['propertyName'],['getetag']]);
		$response = $this->getPropertiesOfFolder(
			$user,
			$path,
			$propertiesTable
		);
		$xmlObject = $this->featureContext->getResponseXml($response);
		$this->storedETAG[$user][$storePath]
			= $this->featureContext->getEtagFromResponseXmlObject($xmlObject);
		return $xmlObject;
	}

	/**
	 * @When user :user stores etag of element :path using the WebDAV API
	 *
	 * @param string $user
	 * @param string $path
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userStoresEtagOfElement(string $user, string $path):void {
		$this->storeEtagOfElement(
			$user,
			$path
		);
	}

	/**
	 * @Given user :user has stored etag of element :path on path :storePath
	 *
	 * @param string $user
	 * @param string $path
	 * @param string $storePath
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userStoresEtagOfElementOnPath(string $user, string $path, string $storePath):void {
		$user = $this->featureContext->getActualUsername($user);
		$this->storeEtagOfElement(
			$user,
			$path,
			$storePath
		);
		if ($storePath == "") {
			$storePath = $path;
		}
		if ($this->storedETAG[$user][$storePath] === null || $this->storedETAG[$user][$path] === "") {
			throw new Exception("Expected stored etag to be some string but found null!");
		}
	}

	/**
	 * @Given user :user has stored etag of element :path
	 *
	 * @param string $user
	 * @param string $path
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userHasStoredEtagOfElement(string $user, string $path):void {
		$user = $this->featureContext->getActualUsername($user);
		$this->storeEtagOfElement(
			$user,
			$path
		);
		if ($this->storedETAG[$user][$path] === "" || $this->storedETAG[$user][$path] === null) {
			throw new Exception("Expected stored etag to be some string but found null!");
		}
	}

	/**
	 * @Then /^the properties response should contain an etag$/
	 *
	 * @return void
	 * @throws Exception
	 */
	public function thePropertiesResponseShouldContainAnEtag():void {
		Assert::assertTrue(
			$this->featureContext->isEtagValid($this->featureContext->getEtagFromResponseXmlObject()),
			__METHOD__
			. " getetag not found in response"
		);
	}

	/**
	 * @Then as user :username the last response should have the following properties
	 *
	 * only supports new DAV version
	 *
	 * @param string $username
	 * @param TableNode $expectedPropTable with following columns:
	 *                                     resource: full path of resource(file/folder/entry) from root of your oc storage
	 *                                     property: expected name of property to be asserted, eg: status, href, customPropName
	 *                                     value: expected value of expected property
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theResponseShouldHavePropertyWithValue(string $username, TableNode $expectedPropTable):void {
		$username = $this->featureContext->getActualUsername($username);
		$this->featureContext->verifyTableNodeColumns($expectedPropTable, ['resource', 'propertyName', 'propertyValue']);
		$responseXmlObject = $this->featureContext->getResponseXml();

		$hrefSplittedUptoUsername = \explode("/", (string)$responseXmlObject->xpath("//d:href")[0]);
		$xmlHrefSplittedArray = \array_slice(
			$hrefSplittedUptoUsername,
			0,
			\array_search($username, $hrefSplittedUptoUsername) + 1
		);
		$xmlHref = \implode("/", $xmlHrefSplittedArray);
		foreach ($expectedPropTable->getColumnsHash() as $col) {
			if ($col["propertyName"] === "status") {
				$xmlPart = $responseXmlObject->xpath(
					"//d:href[.='" .
					$xmlHref . $col["resource"] .
					"']/following-sibling::d:propstat//d:" .
					$col["propertyName"]
				);
			} else {
				$xmlPart = $responseXmlObject->xpath(
					"//d:href[.= '" .
					$xmlHref . $col["resource"] .
					"']/..//oc:" . $col["propertyName"]
				);
			}
			Assert::assertEquals(
				$col["propertyValue"],
				$xmlPart[0],
				__METHOD__
				. " Expected '" . $col["propertyValue"] . "' but got '" . $xmlPart[0] . "'"
			);
		}
	}

	/**
	 * @param string $path
	 * @param string $user
	 *
	 * @return string
	 * @throws Exception
	 */
	public function getCurrentEtagOfElement(string $path, string $user):string {
		$user = $this->featureContext->getActualUsername($user);
		$propertiesTable = new TableNode([['propertyName'],['getetag']]);
		$response = $this->getPropertiesOfFolder(
			$user,
			$path,
			$propertiesTable
		);
		return $this->featureContext->getEtagFromResponseXmlObject($this->featureContext->getResponseXml($response));
	}

	/**
	 * @param string $path
	 * @param string $user
	 * @param string $messageStart
	 *
	 * @return string
	 */
	public function getStoredEtagOfElement(string $path, string $user, string $messageStart = ''):string {
		if ($messageStart === '') {
			$messageStart = __METHOD__;
		}
		Assert::assertArrayHasKey(
			$user,
			$this->storedETAG,
			$messageStart
			. " Trying to check etag of element $path of user $user but the user does not have any stored etags"
		);
		Assert::assertArrayHasKey(
			$path,
			$this->storedETAG[$user],
			$messageStart
			. " Trying to check etag of element $path of user $user but the user does not have a stored etag for the element"
		);
		return $this->storedETAG[$user][$path];
	}

	/**
	 * @Then these etags should not have changed:
	 *
	 * @param TableNode $etagTable
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theseEtagsShouldNotHaveChanged(TableNode $etagTable):void {
		$this->featureContext->verifyTableNodeColumns($etagTable, ["user", "path"]);
		$this->featureContext->verifyTableNodeColumnsCount($etagTable, 2);
		$changedEtagCount = 0;
		$changedEtagMessage = __METHOD__;
		foreach ($etagTable->getColumnsHash() as $row) {
			$user = $row["user"];
			$path = $row["path"];
			$user = $this->featureContext->getActualUsername($user);
			$actualEtag = $this->getCurrentEtagOfElement($path, $user);
			$storedEtag = $this->getStoredEtagOfElement($path, $user, __METHOD__);
			if ($actualEtag !== $storedEtag) {
				$changedEtagCount = $changedEtagCount + 1;
				$changedEtagMessage
					.= "\nThe etag '$storedEtag' of element '$path' of user '$user' changed to '$actualEtag'.";
			}
		}
		Assert::assertEquals(0, $changedEtagCount, $changedEtagMessage);
	}

	/**
	 * @Then these etags should have changed:
	 *
	 * @param TableNode $etagTable
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theseEtagsShouldHaveChanged(TableNode $etagTable):void {
		$this->featureContext->verifyTableNodeColumns($etagTable, ["user", "path"]);
		$this->featureContext->verifyTableNodeColumnsCount($etagTable, 2);
		$unchangedEtagCount = 0;
		$unchangedEtagMessage = __METHOD__;
		foreach ($etagTable->getColumnsHash() as $row) {
			$user = $row["user"];
			$path = $row["path"];
			$user = $this->featureContext->getActualUsername($user);
			$actualEtag = $this->getCurrentEtagOfElement($path, $user);
			$storedEtag = $this->getStoredEtagOfElement($path, $user, __METHOD__);
			if ($actualEtag === $storedEtag) {
				$unchangedEtagCount = $unchangedEtagCount + 1;
				$unchangedEtagMessage
					.= "\nThe etag '$storedEtag' of element '$path' of user '$user' did not change.";
			}
		}
		Assert::assertEquals(0, $unchangedEtagCount, $unchangedEtagMessage);
	}

	/**
	 * @Then /^the etag of element "([^"]*)" of user "([^"]*)" (should|should not) have changed$/
	 *
	 * @param string $path
	 * @param string $user
	 * @param string $shouldShouldNot
	 *
	 * @return void
	 * @throws Exception
	 */
	public function etagOfElementOfUserShouldOrShouldNotHaveChanged(string $path, string $user, string $shouldShouldNot):void {
		$user = $this->featureContext->getActualUsername($user);
		$actualEtag = $this->getCurrentEtagOfElement($path, $user);
		$storedEtag = $this->getStoredEtagOfElement($path, $user, __METHOD__);
		if ($shouldShouldNot === 'should not') {
			Assert::assertEquals(
				$storedEtag,
				$actualEtag,
				__METHOD__
				. " The etag of element '$path' of user '$user' was not expected to change."
				. " The stored etag was '$storedEtag' but got '$actualEtag' from the response"
			);
		} else {
			Assert::assertNotEquals(
				$storedEtag,
				$actualEtag,
				__METHOD__
				. " The etag of element '$path' of user '$user' was expected to change."
				. " The stored etag was '$storedEtag' and also got '$actualEtag' from the response"
			);
		}
	}

	/**
	 * This will run before EVERY scenario.
	 * It will set the properties for this object.
	 *
	 * @BeforeScenario
	 *
	 * @param BeforeScenarioScope $scope
	 *
	 * @return void
	 */
	public function before(BeforeScenarioScope $scope):void {
		// Get the environment
		$environment = $scope->getEnvironment();
		// Get all the contexts you need in this context
		$this->featureContext = $environment->getContext('FeatureContext');
	}
}
