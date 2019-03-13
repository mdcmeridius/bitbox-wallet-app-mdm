/**
 * Copyright 2018 Shift Devices AG
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, h, RenderableProps } from 'preact';
import { CopyableInput } from '../../../components/copy/Copy';
import { Button } from '../../../components/forms';
import { QRCode } from '../../../components/qrcode/qrcode';
import { translate, TranslateProps } from '../../../decorators/translate';
import { apiPost } from '../../../utils/request';
import { AccountInfo } from '../account';

interface ProvidedProps {
    info: AccountInfo;
    code: string;
}

type Props = ProvidedProps & TranslateProps;

class SigningConfiguration extends Component<Props> {

    private verifyPublicKey = (index: number) => {
        apiPost(`account/${this.props.code}/verify-publickey`, index);
    }

    public render({ t, info }: RenderableProps<Props>) {
        return (
        // TODO: add info if single or multisig, and threshold.
        <div>
            {
                info.signingConfiguration.xpubs.map((xpub, index) => {
                    return (
                        <div key={xpub}>
                            <strong>
                                {t('accountInfo.extendedPublicKey')}
                                {info.signingConfiguration.xpubs.length > 1 && (' #' + (index + 1))}
                            </strong><br />
                            <QRCode data={xpub} />
                            <CopyableInput value={xpub} />
                            <Button primary onClick={() => this.verifyPublicKey(index)}>
                                {t('Verify')}
                            </Button>
                        </div>
                    );
                })
            }
        </div>
    ); }
}

const HOC = translate<ProvidedProps>()(SigningConfiguration);
export { HOC as SigningConfiguration };
